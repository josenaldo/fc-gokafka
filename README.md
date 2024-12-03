# FC-GoKafa

O Kafka é um sistema de mensageria distribuído que permite a troca de mensagens entre aplicações. Ele é muito utilizado para a criação de sistemas de eventos, como o Event Sourcing.

- [FC-GoKafa](#fc-gokafa)
  - [Descrição](#descrição)
  - [Requisitos](#requisitos)
  - [Referências](#referências)
    - [Apache Kafka](#apache-kafka)
    - [Docker](#docker)
    - [Go](#go)
    - [librdkafka](#librdkafka)
  - [Comandos](#comandos)
    - [Container](#container)
      - [Iniciar Containers](#iniciar-containers)
      - [Parar o container](#parar-o-container)
      - [Conectar ao Container](#conectar-ao-container)
    - [Comandos do Kafka](#comandos-do-kafka)
      - [Tópicos](#tópicos)
        - [Criar um tópico](#criar-um-tópico)
        - [Listar tópicos](#listar-tópicos)
        - [Descrever um tópico](#descrever-um-tópico)
        - [Deletar um tópico](#deletar-um-tópico)
      - [Consumir mensagens](#consumir-mensagens)
        - [Consumir mensagens de um tópico](#consumir-mensagens-de-um-tópico)
        - [Consumir mensagens a partir do início](#consumir-mensagens-a-partir-do-início)
        - [Consumir mensagens de um grupo de consumidores](#consumir-mensagens-de-um-grupo-de-consumidores)
      - [Produzir mensagens](#produzir-mensagens)
      - [Analisando os grupos de consumidores](#analisando-os-grupos-de-consumidores)

## Descrição

Esse repositório contém um exemplo de como usar o Kafka com Docker. Ele foi criado durante o curso 'Apache Kafka', no 'Curso FullCycle 3.0', da [Full Cycle](https://curso.fullcycle.com.br/curso-fullcycle-3-0/).

## Requisitos

- Docker
- Docker Compose
- Go

## Referências

### Apache Kafka

Apache Kafka é uma plataforma de streaming distribuída. É usada para construir pipelines de dados em tempo real e aplicações de streaming. É escalável horizontalmente, tolerante a falhas e extremamente rápido.

- [Apache Kafka](https://kafka.apache.org/)

### Docker

O docker é uma plataforma de software que permite a criação, o teste e a implantação de aplicativos rapidamente. Ele é projetado para fornecer uma maneira fácil e leve de criar contêineres, que são ambientes isolados que podem ser executados em qualquer lugar.

- [Docker](https://www.docker.com/)

### Go

Go é uma linguagem de programação de código aberto que facilita a criação de software simples, confiável e eficiente.

- [Go](https://golang.org/)

### librdkafka

Librdkafka é uma implementação da biblioteca C/C++ do protocolo Apache Kafka, contendo suporte para Produtor e Consumidor. Foi projetado com confiabilidade na entrega de mensagens e alto desempenho em mente, com números atuais superiores a 1 milhão de mensagens/segundo para o produtor e 3 milhões de mensagens/segundo para o consumidor.

- [librdkafka](https://github.com/confluentinc/librdkafka)

## Comandos

Para inciar, parar e se conectar ao conteiner do Kafka, você pode usar os comandos abaixo. Uma outra opção é usar o arquivo `Makefile` para executar esses comandos.

### Container

O docker-container.yaml contém as configurações para os seguintes containers:

- app: Container com o Go instalado.
- kafka: Container com o Kafka instalado.
- zookeeper: Container com o Zookeeper instalado.
- conteol-center: Container com o Control Center instalado.
- kafka-ui: Container com o Kafka UI instalado.

#### Iniciar Containers

Para iniciar os containers, execute o comando abaixo.

```bash
docker compose up -d --build
```

Se quiser usar o arquivo `Makefile`, execute o comando abaixo.

```bash
make up
```

#### Parar o container

Para parar e remover os containers, execute o comando abaixo.

```bash
docker compose down
```

Se quiser usar o arquivo `Makefile`, execute o comando abaixo.

```bash
make down-kafka
```

#### Conectar ao Container

Para conectar aos containers, execute os comandos abaixo, de acordo com o container:

```bash
docker exec -it gokafka bash # Container com o Go instalado
docker exec -it kafka bash # Container com o Kafka instalado
```

Se quiser usar o arquivo `Makefile`, execute o comando abaixo.

```bash
make exec-app # Container com o Go instalado
make exec-kafka # Container com o Kafka instalado
```

### Comandos do Kafka

Após estar conectado ao container do Kafka, é possível executar comandos do Kafka. Abaixo estão alguns exemplos.

Observe que todos os comandos contém a flag `--bootstrap-server localhost:9092`, que indica o endereço do servidor Kafka. Essa flag éobrigatória para todos os comandos.

#### Tópicos

##### Criar um tópico

```bash
kafka-topics --bootstrap-server localhost:9092 --create --topic teste --partitions 3 --replication-factor 1
```

##### Listar tópicos

```bash
kafka-topics --bootstrap-server localhost:9092 --list
```

##### Descrever um tópico

```bash
kafka-topics --bootstrap-server localhost:9092 --describe --topic teste
```

##### Deletar um tópico

```bash
kafka-topics --bootstrap-server localhost:9092 --delete --topic teste
```

#### Consumir mensagens

##### Consumir mensagens de um tópico

Para consumir mensagens de um tópico, execute o comando ´kafka-console-consumer´.

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic teste
```

##### Consumir mensagens a partir do início

Para consumir mensagens a partir do início, use a flag `--from-beginning`.

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic teste --from-beginning
```

##### Consumir mensagens de um grupo de consumidores

Para consumir mensagens de um grupo de consumidores, use a flag `--group`.

Se dois consumidores estiverem em grupos diferentes, cada um receberá todas as mensagens.

Se mais de um consumidor estiver no mesmo grupo, as mensagens serão distribuídas entre eles.

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic teste --group x
```

#### Produzir mensagens

Para produzir mensagens em um tópico, execute o comando ´kafka-console-producer´.

```bash
kafka-console-producer --bootstrap-server localhost:9092 --topic teste
```

#### Analisando os grupos de consumidores

Para ver os grupos de consumidores, execute o comando abaixo.

```bash
kafka-consumer-groups --bootstrap-server localhost:9092 --list
```

Para ver os detalhes de um grupo de consumidores, execute o comando abaixo.

```bash
kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group x
```


