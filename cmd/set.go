package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		setConfig(args[0])
	},
}

func setConfig(configPath string) {
	cc.Path = configPath
	bytes, err := json.Marshal(cc)
	if err != nil {
		log.Fatalln("json marshal path: ["+configPath+"] failed", err)
	}
	os.WriteFile(ConfigPath(), bytes, os.ModePerm)
}
