build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

tidy:
	docker-compose run --rm api go mod tidy

mysql:
	docker-compose exec db bash -c "mysql -u root -ppassword"
