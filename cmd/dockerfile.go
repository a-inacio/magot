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
	"github.com/a-inacio/magot/internal/template"
	"github.com/a-inacio/magot/internal/template/dockerfile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dockerfileCmd represents the dockerfile command
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generate a dockerfile",
	Long:  `Generate a dockerfile for your application`,
	Run: func(cmd *cobra.Command, args []string) {
		dockerfile := dockerfile.NewDockerfile(viper.GetBool("useMakefile"))

		var err error
		if viper.GetBool("preview") {
			err = template.Preview(dockerfile)
		} else {
			err = template.WriteFile(dockerfile, "Dockerfile")
		}

		if err != nil {
			panic("Something went really wrong! Bailing out...")
		}
	},
}

func init() {
	dockerfileCmd.PersistentFlags().Bool("makefile", false, "use a makefile to build the application")
	cobra.CheckErr(viper.BindPFlag("useMakefile", dockerfileCmd.PersistentFlags().Lookup("makefile")))

	rootCmd.AddCommand(dockerfileCmd)
}
