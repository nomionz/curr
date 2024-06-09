FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o converter .

FROM alpine:3.18

WORKDIR /app

# html templates
COPY --from=builder /app/views ./views

# timezone info
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

# binary
COPY --from=builder /app/converter .

EXPOSE 80

CMD ["./converter"]
