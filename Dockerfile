# Base build image
FROM golang:1.12-alpine AS build_base

# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /home/digidoc

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

# This image builds the digidoc server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .

# And compile the project
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS digidoc
# We add the certificates to be able to verify remote digidoc instances
RUN apk add ca-certificates curl
# Finally we copy the statically compiled Go binary.
COPY --from=server_builder /go/bin/digidoc ./build/digidoc
COPY ./email/templates ./email/templates
COPY ./location/data ./location/data

EXPOSE 8080

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f ${DIGIDOC_DB_HOST}:${DIGIDOC_DB_PORT} || exit 1

ENTRYPOINT ["./build/digidoc server start"]
