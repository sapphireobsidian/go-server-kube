FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/demo-server-blank

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/demo-server-blank .

# Start fresh from a smaller image
FROM alpine:3.13
RUN apk add ca-certificates

COPY --from=build_base /tmp/demo-server-blank/out/demo-server-blank /app/demo-server-blank

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/demo-server-blank"]