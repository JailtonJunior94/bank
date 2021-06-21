FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o loan loan/loan.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/loan .
COPY ["loan/appsettings.Development.json","."]
COPY ["loan/appsettings.Production.json","."]
COPY ["loan/appsettings.Staging.json","."]

EXPOSE 9000

CMD [ "/app/loan" ]