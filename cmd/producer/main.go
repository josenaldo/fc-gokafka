package main

import (
	"fmt"
	"log"

	"math/rand"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event)

	producer := NewKafkaProducer()

	key := fmt.Sprintf("%d", rand.Intn(100))
	message := fmt.Sprintf("Mensagem de teste %s", key)

	err := Publish(message, "teste", producer, nil, deliveryChannel)
	if err != nil {
		panic(err)
	}
	go DeliverReport(deliveryChannel)

	fmt.Println("Mensagem enviada com sucesso")
	producer.Flush(1000)

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		// Endereço do servidor kafka
		"bootstrap.servers": "kafka:9092",

		// Tempo máximo de espera para entrega da mensagem. 0 = entrega síncrona, -1 = entrega assíncrona
		"delivery.timeout.ms": "0",

		// Número de replicas que devem confirmar o recebimento da mensagem. 1 = apenas o líder, all = todas as replicas
		"acks": "all",

		//When set to true, the producer will ensure that messages are successfully produced exactly once and in the original produce order. The following configuration properties are adjusted automatically (if not modified by the user) when idempotence is enabled: max.in.flight.requests.per.connection=5 (must be less than or equal to 5), retries=INT32_MAX (must be greater than 0), acks=all, queuing.strategy=fifo. Producer instantation will fail if user-supplied configuration is incompatible.

		// Quando definido como true, o produtor garantirá que as mensagens sejam produzidas com exatidão e na ordem
		// original de produção. As seguintes propriedades de configuração são ajustadas automaticamente (se não forem
		// modificadas pelo usuário) quando a idempotência é ativada:
		// - max.in.flight.requests.per.connection=5 (deve ser menor ou igual a 5),
		// - retries=INT32_MAX (deve ser maior que 0),
		// - acks=all,
		// - queuing.strategy=fifo.
		// A instanciação do produtor falhará se a configuração fornecida pelo usuário for incompatível.
		"enable.idempotence": "true",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println("Error producer: ", err.Error())
	}

	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, deliveryChannel)
	if err != nil {
		return err
	}

	return nil
}

func DeliverReport(deliveryChannel chan kafka.Event) {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Println("Error sending message: ", ev.TopicPartition.Error)
			} else {
				log.Println("Message sent: ", ev.TopicPartition)
			}
		}
	}
}
