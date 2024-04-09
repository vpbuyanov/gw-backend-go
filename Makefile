create:
	migrate create -ext sql -dir migrations -seq $(name)

up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

deploy:
	docker compose up -d

dev:
	docker build -t vpbuyanov/gw-backend-go:latest .
	docker-compose up -d