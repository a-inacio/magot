#
# Build layer
#

FROM golang:1.17.2-alpine3.14 AS build-env

RUN apk add --update make protoc protobuf protobuf-dev git build-base bash curl

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u google.golang.org/grpc
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1

# Copy everything of current project
WORKDIR /src

COPY . /src/
ENV GO111MODULE="on" \
    CGO_ENABLED=1 \
    GOOS=linux


RUN make build


#
# Runtime layer
#

FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /src/main .

ENTRYPOINT ["/app/main"]
