FROM golang:latest

COPY ./static/ /static
COPY ./migrations/ /migrations
COPY ./bin/boats .
COPY ./boats.toml .

CMD ["./boats", "-conf=boats.toml"]
EXPOSE 9876
