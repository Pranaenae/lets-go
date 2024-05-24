# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-binary ./cmd/web

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-binary /go-binary

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-binary"]