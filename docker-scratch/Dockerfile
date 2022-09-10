# syntax=docker/dockerfile:1

## Build
# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY *.go ./

RUN go build -o /main

RUN adduser \
    -h "/dev/null" \
    -g "" \
    -s "/sbin/nologin" \
    -D \
    -H \
    -u 10001 \
    playerone

## Deploy
FROM scratch

WORKDIR /


COPY --from=0 /main /main

COPY --from=0 /etc/passwd /etc/passwd

# USER 10001 works. However, the gid=0(root). So, don't do USER 10001.
USER 10001:10001

# ADD busybox.tar.xz /
# CMD [ "sh" ]

ENTRYPOINT [ "/main" ]