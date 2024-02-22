FROM golang:1.21.4-alpine3.18 AS builder

RUN mkdir -p /app
WORKDIR /app


COPY .env ./
COPY go.mod ./

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0
RUN export GOOS=linux
RUN go build -o ./ /app/cmd/main.go
RUN mv main /

FROM alpine

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/.env /
COPY --from=builder /main /

EXPOSE 7053

CMD ["/main"]