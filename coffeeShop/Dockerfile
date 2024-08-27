FROM golang:1.22.3 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o coffeeshop_app .

# FROM golang:1.22.3 // stage two
# The minimal docker stage
FROM scratch
COPY --from=builder /app ./

EXPOSE 8080

CMD ["./coffeeshop_app"]