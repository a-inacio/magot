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

import "github.com/a-inacio/magot/internal/template"

// Dockerfile abstracts the recipe to render dockerfiles for specific languages<F5>
type Dockerfile struct {
	base        template.Template
	useMakefile bool
}

func (d Dockerfile) Source() string {
	return d.base.Source()
}

func (d Dockerfile) Model() interface{} {
	return d.base.Model()
}

func NewDockerfile(useMakefile bool, buildLayer string, runtimeLayer string) *Dockerfile {
	return &Dockerfile{NewGoDockerfile(useMakefile, "\"/app/main\"", buildLayer, runtimeLayer), useMakefile}
}
