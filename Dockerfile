FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o pdns_filler ./cmd/pdns_filler

FROM scratch
COPY --from=builder /app/pdns_filler /pdns_filler
ENTRYPOINT [ "/pdns_filler" ]
