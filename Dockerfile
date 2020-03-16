FROM golang:1.13.0-stretch AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN make build

FROM scratch

WORKDIR /root

COPY --from=builder /app/bin .

EXPOSE 3000

ENTRYPOINT ["./summer"]
