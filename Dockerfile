FROM golang:alpine as builder
RUN apk update
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go clean
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/api/
RUN CGO_ENABLED=0 GOOS=linux go test -v ./...
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]