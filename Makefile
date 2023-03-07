start:
	docker-compose up

rebuild:
	docker system prune
	docker-compose up --build

stop:
	docker-compose down

console:
	docker exec -ti rss_reader bash