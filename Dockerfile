FROM golang:1.16-alpine AS build

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go build -o /app

FROM alpine:latest

WORKDIR /

COPY --from=build /app ./

EXPOSE 8080

RUN addgroup -S app && adduser -S app -G app
USER app:app

CMD ["/app"]
