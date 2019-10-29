FROM golang:alpine

COPY ./static/ /static
COPY ./bin/boats .
COPY ./boats.toml .

CMD ["./boats", "-conf=boats.toml"]
