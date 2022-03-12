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
package compose

import (
	"errors"
)

// Compose abstracts the recipe to render docker compose files
type Compose struct {
	UseApp      bool
	UsePostgres bool
}

func (c Compose) Source() string {
	return string(dockercomposeTemplate())
}

func (c Compose) Model() interface{} {
	return c
}

func (c Compose) IsValid() error {
	if c.UseApp || c.UsePostgres {
		return nil
	}

	return errors.New("At least one service needs to be enabled!")
}

func NewCompose(useApp bool, usePostgres bool) *Compose {
	return &Compose{useApp, usePostgres}
}

func dockercomposeTemplate() []byte {
	return []byte(`#
# WIP
#
services:
{{ if .UseApp }} app:
    build: . {{ end }}
`)
}
