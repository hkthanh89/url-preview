ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -mod=readonly -v -o /run-app .


FROM debian:bookworm

RUN apt update && apt install -y ca-certificates

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
