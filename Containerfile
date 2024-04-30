FROM docker.io/golang:1.22.2-alpine3.19 AS builder

WORKDIR /app
ENV CGO_ENABLED=1

RUN apk -U --no-cache add git gcc musl-dev sqlite
COPY . .
RUN go build
RUN sqlite3 deities.sqlite ".read extras/deities.sql"


FROM docker.io/alpine:3.19

LABEL LastUpdate="2024/04/30-4"
COPY --from=builder /app/wp24-deities /wp24-deities
COPY --from=builder /app/deities.sqlite /deities.sqlite
RUN apk -U --no-cache upgrade

EXPOSE 8080

ENTRYPOINT ["/wp24-deities"]
