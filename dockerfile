FROM golang:1.22.5-alpine3.20 AS builder

COPY . /github.com/Naumovets/go-auth/source/
WORKDIR /github.com/Naumovets/go-auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Naumovets/go-auth/source/bin/auth_server .

CMD [ "./auth_server" ]
