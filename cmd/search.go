package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Outf("%s", viper.AllKeys)
	},
}
