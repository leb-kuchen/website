# Build.
FROM golang:1.22.1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint cmd/api/main.go

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/cmd/web/css /app/cmd/web/js /app/cmd/web/
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
