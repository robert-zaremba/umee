# Builder docker container
FROM golang:1.19-alpine AS builder
RUN apk add --no-cache make git libc-dev gcc linux-headers build-base && \
    apk add --update ca-certificates

WORKDIR /src/app/
COPY . .

# Cosmwasm - Download correct libwasmvm version
# RUN WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | cut -d ' ' -f 2) && \
#     wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$(uname -m).a \
#       -O /lib/libwasmvm_muslc.a && \
#     # verify checksum
#     wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
#     sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep $(uname -m) | cut -d ' ' -f 1)

# Build the binary
RUN LEDGER_ENABLED=false make install

# Final image
FROM alpine:edge
RUN apk add bash curl jq && \
    apk add --update ca-certificates

COPY --from=builder /go/bin/umeed /usr/local/bin/
EXPOSE 26656 26657 1317 9090

# Run umeed by default, omit entrypoint to ease using container with CLI
CMD ["umeed"]
STOPSIGNAL SIGTERM
