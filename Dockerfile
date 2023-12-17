FROM golang:alpine

ENV HTTPPORT=6789
ENV PGDSN=

RUN mkdir /app
ADD cmd /app/cmd
ADD internal /app/internal
COPY go.* /app/
COPY --chmod=555 docker-entrypoint.sh /app/
WORKDIR /app

RUN go mod download && go build -o main cmd/main.go
ENTRYPOINT ["/app/docker-entrypoint.sh"]
