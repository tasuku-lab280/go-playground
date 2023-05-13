include .env

build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

tidy:
	docker-compose exec api go mod tidy

mysql:
	docker-compose exec db bash -c "mysql -u ${DATABASE_USERNAME} -p${DATABASE_PASSWORD}"

db/reinstall:
	make migrate/down && make migrate/up

migrate/create:
	docker-compose exec api migrate create -ext sql -dir db/migrate ${NAME}

migrate/%:
	docker compose exec api migrate \
		-database="mysql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@tcp(${DATABASE_ENDPOINT}:${DATABASE_PORT})/${DATABASE_NAME}" \
		-path="db/migrate" \
		$(subst migrate/,,$@) ${N}
