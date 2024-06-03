FROM golang:1.22 as builder

WORKDIR /build

COPY go.mod go.sum ./
COPY ./src ./src

RUN go mod download
RUN ls -ll ./src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /build/main .

EXPOSE 8080

CMD ["/app/main"]