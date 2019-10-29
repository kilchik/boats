build:
	GOOS=linux go build -o ./bin/boats ./cmd/boats/main.go
	docker build -t boats .
