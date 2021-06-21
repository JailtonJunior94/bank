FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o risk risk/risk.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/risk .
COPY ["risk/settings.Development.json","."]
COPY ["risk/settings.Production.json","."]
COPY ["risk/settings.Staging.json","."]

CMD [ "/app/risk" ]