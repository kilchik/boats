run:
	GOOS=linux go build -o ./bin/boats ./cmd/boats/main.go
	docker-compose build --no-cache
	docker-compose up
