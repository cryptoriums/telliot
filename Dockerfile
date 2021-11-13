# When this issue is fixed can upgrade to latest https://github.com/golang/go/issues/49366
FROM golang:1.17.2-buster AS builder
WORKDIR /go/src
RUN apt-get update
COPY ./ .
RUN make build

FROM debian:buster-slim  
WORKDIR /
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /go/src/telliot .
ENTRYPOINT ["./telliot"]