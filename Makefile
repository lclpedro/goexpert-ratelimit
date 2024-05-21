test:
	docker-compose up -d
	go test -v ./...
	docker rm -f redis-rateliter app-rateliter