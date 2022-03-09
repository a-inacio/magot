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
	"io"
	"os"
	tpl "text/template"
)

func Preview(t Template) error {
	return write(t, os.Stdout)
}

func WriteFile(t Template, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return write(t, file)
}

func write(t Template, wr io.Writer) error {

	fileTemplate, err := tpl.New("tpl").Parse(t.Source())
	if err != nil {
		return err
	}

	err = fileTemplate.Execute(wr, t.Model())
	//err = fileTemplate.Execute(file, t.Model())
	if err != nil {
		return err
	}

	return nil
}
