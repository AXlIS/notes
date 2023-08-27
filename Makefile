run:
	docker-compose up --build

.DEFAULT_GOAL := run

migrate:
	docker-compose exec app migrate -path ./migrations -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up