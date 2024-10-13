FROM golang:1.23-alpine as golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tmp/main .

FROM gcr.io/distroless/static-debian11

COPY --from=golang /app/tmp/main server

EXPOSE 8080

CMD ["./server"]
