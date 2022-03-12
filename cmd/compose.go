/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/a-inacio/magot/internal/template"
	"github.com/a-inacio/magot/internal/template/compose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Generate a dockercompose file",
	Long: `Generate a dockercompose file with a curated list of practical services.

	The list comes from common components that you may require to test your application.
	Don't expect the parameters to be ready for production use!

	Check the help for this command for more details about what you can get.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		compose := compose.NewCompose(
			viper.GetBool("useApp"),
			viper.GetBool("usePostgres"),
		)

		var err error

		err = compose.IsValid()

		if err != nil {
			return err
		}

		if viper.GetBool("preview") {
			err = template.Preview(compose)
		} else {
			err = template.WriteFile(compose, "docker-compose.yml")
		}

		return err
	},
}

func init() {
	composeCmd.PersistentFlags().Bool("app", false, "use a generic service for the current application dockerfile")
	cobra.CheckErr(viper.BindPFlag("useApp", composeCmd.PersistentFlags().Lookup("app")))

	composeCmd.PersistentFlags().Bool("postgres", false, "use a simple postgres instance")
	cobra.CheckErr(viper.BindPFlag("usePostgres", composeCmd.PersistentFlags().Lookup("postgres")))

	rootCmd.AddCommand(composeCmd)
}
