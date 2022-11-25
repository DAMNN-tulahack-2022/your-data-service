FROM golang:1.18.3-alpine3.16 as build

WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o adm ./cmd/supadm_service/main.go

FROM alpine:3.16 as supadm
RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add ca-certificates --no-cache

WORKDIR /usr/local/app
COPY --from=build /usr/local/app/adm bin/adm

CMD bin/adm