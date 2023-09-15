FROM golang:alpine as builder

WORKDIR /src/app
COPY . .

RUN go build -o backend

FROM alpine

WORKDIR /src/app
COPY --from=builder /src/app/backend .

EXPOSE 8000
CMD ["./backend"]
