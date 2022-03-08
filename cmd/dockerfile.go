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
package cmd

import (
	"fmt"

	"github.com/a-inacio/magot/internal/template"
	"github.com/spf13/cobra"
)

// dockerfileCmd represents the dockerfile command
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generate a dockerfile",
	Long:  `Generate a dockerfile for your application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(string(template.DockerfileTemplate()))
	},
}

func init() {
	rootCmd.AddCommand(dockerfileCmd)
}
