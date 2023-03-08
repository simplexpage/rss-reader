start:
	docker-compose up

rebuild:
	docker system prune
	docker-compose up --build

stop:
	docker-compose down

console:
	docker exec -ti rss_reader bash

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

race:
	go test -v -race -count=1 ./...

test:
	go test -v -count=1 ./...

