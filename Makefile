.PHONY: create
create:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: up
up:
	migrate -database postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable -path migrations up

.PHONY: down
down:
	migrate -database postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable -path migrations down

.PHONY: dev
dev:
	docker build -t vpbuyanov/gw-backend-go:latest .
	docker-compose up -d
	docker stop gw-backend-go
	docker cp config.yaml gw-backend-go:/
	docker start gw-backend-go