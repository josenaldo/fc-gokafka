up:
	docker compose up -d --build

down:
	docker compose down

exec-app:
	docker exec -it gokafka bash

exec-kafka:
	docker exec -it kafka bash

run:
	go run cmd/producer/main.go

.PHONY: up down exec-app exec-kafka run