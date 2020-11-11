FROM golang:1.12.0-alpine3.9 AS builder
WORKDIR /go/src/github.com/call-me-snake/date_service
COPY . .
RUN go install ./...

FROM golang:1.12.0-alpine3.9 AS production
COPY --from=builder /go/bin/cmd ./app
