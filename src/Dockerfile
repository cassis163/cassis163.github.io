# Build.
FROM golang:1.21.2 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN $GO_PATH/bin/templ generate .
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=build-stage /entrypoint /entrypoint
COPY --chown=nonroot --from=build-stage /app/assets /assets
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
