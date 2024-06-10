create:
	migrate create -ext sql -dir migrations -seq $(name)

up:
	migrate -database postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable -path migrations up

down:
	migrate -database postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable -path migrations down

deploy:
	docker compose up -d

dev:
	docker build -t vpbuyanov/gw-backend-go:latest .
	docker-compose up -d
	docker stop gw-backend-go
	docker cp config.yaml gw-backend-go:/
	docker start gw-backend-go