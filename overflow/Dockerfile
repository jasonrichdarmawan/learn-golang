FROM golang:1.19-alpine AS stage1

COPY main.go ./
RUN GO111MODULE=off go build -o /main

FROM scratch

COPY --from=0 /lib/ld-musl-aarch64.so.1 /lib/ld-musl-aarch64.so.1
COPY --from=stage1 /main /main

ENTRYPOINT [ "/main" ]