FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN mkdir -p /app/logs
RUN touch /app/logs/errorLogs.txt
RUN ls -la

COPY ./web/html /app/web/html
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./acmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
