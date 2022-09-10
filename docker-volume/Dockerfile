# syntax=docker/dockerfile:1
FROM alpine:3.16

RUN mkdir /secret-dir && chown 10001:10001 /secret-dir && chmod 750 /secret-dir

RUN adduser \
    -h "/dev/null" \
    -g "" \
    -s "/sbin/nologin" \
    -D \
    -H \
    -u 10001 \
    playerone

USER playerone:playerone