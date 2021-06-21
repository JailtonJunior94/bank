FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o customer customer/customer.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/customer .
COPY ["customer/settings.Development.json","."]
COPY ["customer/settings.Production.json","."]
COPY ["customer/settings.Staging.json","."]

EXPOSE 8000

CMD [ "/app/customer" ]