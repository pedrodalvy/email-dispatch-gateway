FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server ./cmd/api/main.go

FROM scratch

WORKDIR /app
COPY --from=buiLder /app/server server
COPY --from=buiLder /app/.env .env

CMD [ "./server" ]
