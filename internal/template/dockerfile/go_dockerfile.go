/*
Copyright © 2022 António Inácio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package dockerfile

type GoDockerfile struct {
	EntryPointArgs string
}

func NewGoDockerfile() *GoDockerfile {
	return &GoDockerfile{}
}

func (d GoDockerfile) Render() string {
	return string(GoDockerfileTemplate())
}

func GoDockerfileTemplate() []byte {
	return []byte(`#
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

RUN go mod tidy -v && go build -o main

#
# Runtime layer
#

FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /src/main .

ENTRYPOINT ["/app/main"]
`)
}
