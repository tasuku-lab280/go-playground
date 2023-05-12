build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

tidy:
	docker-compose run --rm api go mod tidy
