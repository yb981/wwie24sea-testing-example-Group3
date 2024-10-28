package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "testing",
	Short: "Hugo is a very fast static site generator",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
