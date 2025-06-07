FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git curl sqlite

RUN go install github.com/cosmtrek/air@v1.49.0
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3

COPY api/go.mod api/go.sum ./

RUN go mod download
RUN go mod tidy

COPY api/*.go ./
COPY api/.air.toml ./

EXPOSE 8181

CMD ["/go/bin/air"]
