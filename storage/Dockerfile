FROM golang:1.19.3-alpine AS builder
RUN adduser -D -g '' elf
WORKDIR /opt/app/
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/storage ./cmd/app

FROM alpine:3.17.3
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=elf:1000 /go/bin/storage /storage
COPY --from=builder --chown=elf:1000 /opt/app/.env /.env
USER elf
ENTRYPOINT ["./storage"]