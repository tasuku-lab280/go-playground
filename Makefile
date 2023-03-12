build:
	docker-compose build

up:
	docker-compose up

up-d:
	docker-compose up -d

down:
	docker-compose down

ps:
	docker ps

run:
	@read -p "Enter additional command: " input; \
	docker-compose run api $$input

mysql:
	docker exec -it go_playground_db mysql -u root -p
