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

// GoDockerfile is a Dockerfile Renderer for Go applications
type GoDockerfile struct {
	UseMakefile       bool
	Entrypoint        string
	BuildLayerImage   string
	RuntimeLayerImage string
}

func NewGoDockerfile(
	useMakefile bool,
	entrypoint string,
	buildLayerImage string,
	runtimeLayerImage string) *GoDockerfile {
	return &GoDockerfile{
		useMakefile,
		entrypoint,
		buildLayerImage,
		runtimeLayerImage}
}

func (d GoDockerfile) Source() string {
	return string(goDockerfileTemplate())
}

func (d GoDockerfile) Model() interface{} {
	return d
}

func (d GoDockerfile) IsValid() error {
	return nil
}

func goDockerfileTemplate() []byte {
	return []byte(`#
# Build layer
#

FROM {{.BuildLayerImage}} AS build-env

RUN apk add --update make protoc protobuf protobuf-dev git build-base bash curl

RUN go get -u google.golang.org/protobuf
RUN go get -u google.golang.org/grpc
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1

# Copy everything of current project
WORKDIR /src

COPY . /src/
ENV GO111MODULE="on" \
    CGO_ENABLED=1 \
    GOOS=linux

{{ if .UseMakefile }}
RUN make build
{{ else }}
RUN go mod tidy -v && go build -o main
{{ end }}

#
# Runtime layer
#

FROM {{.RuntimeLayerImage}}
WORKDIR /app
COPY --from=build-env /src/main .

ENTRYPOINT [{{.Entrypoint}}]
`)
}
