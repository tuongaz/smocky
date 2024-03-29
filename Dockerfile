# syntax = docker/dockerfile:1.2

FROM golang:1.18.4-alpine AS build
WORKDIR /src
RUN apk add --no-cache file git
ENV GOMODCACHE /root/.cache/gocache
RUN --mount=target=. --mount=target=/root/.cache,type=cache \
    CGO_ENABLED=0 go build -o /out/smocky -ldflags '-s -d -w' ./; \
    file /out/smocky | grep "statically linked"

FROM scratch
COPY --from=build /out/smocky /bin/smocky
ENTRYPOINT ["/bin/smocky"]
