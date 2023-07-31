container=app

up:
	docker-compose up -d

build:
	docker-compose rm -vsf
	docker-compose down -v --remove-orphans
	docker system prune -a
	docker-compose build
	docker-compose up -d

down:
	docker-compose down
tail-logs:
	docker-compose logs -f ${container}

jumpin:
	docker-compose exec ${container} bash