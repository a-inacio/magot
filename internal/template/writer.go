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
package template

import (
	"os"
	tpl "text/template"
)

func write(renderer Renderer, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileTemplate, err := tpl.New("tpl").Parse(renderer.Render())
	if err != nil {
		return err
	}

	err = fileTemplate.Execute(file, renderer.Data())
	if err != nil {
		return err
	}

	return nil
}
