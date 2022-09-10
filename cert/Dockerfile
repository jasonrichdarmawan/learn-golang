FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./

RUN go build -o /main ./main.go

FROM alpine:3.16

WORKDIR /

COPY --from=0 /main /main

ENTRYPOINT ["/main"]