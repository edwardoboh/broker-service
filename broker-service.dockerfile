FROM golang:1.19-alpine as build

RUN mkdir /app

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api/

RUN chmod +x /app/brokerApp

FROM alpine as production

RUN mkdir /app

WORKDIR /app

COPY --from=build /app/brokerApp .

CMD ["./brokerApp"]