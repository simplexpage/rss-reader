start_local:
	docker-compose up

stop_local:
	docker-compose down

start_local_rebuild:
	docker system prune
	docker-compose up --build

console:
	docker exec -ti rss_reader bash