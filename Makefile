include settings.env

build:
	docker-compose build
run:
	docker-compose up -d
stop:
	docker-compose stop
migrate-up:
	migrate -path ./internal/schema -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up
migrate-down:
	migrate -path ./internal/schema -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" down
logs:
	docker-compose logs -f
status:
	docker-compose ps