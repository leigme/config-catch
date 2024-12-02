package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/leigme/config-catch/cmd/model"
	"github.com/spf13/viper"
)

var cc model.Config = model.Config{}

func init() {
	bytes, err := os.ReadFile(ConfigPath())
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("read config is failed", err)
		}
		bytes, err = json.Marshal(cc)
		err := os.WriteFile(ConfigPath(), bytes, os.ModePerm)
		if err != nil {
			log.Fatalln("create config is failed", err)
		}
	}
	err = json.Unmarshal(bytes, &cc)
	if err != nil {
		log.Fatalln("read config to json failed", err)
	}
	var configType string
	if strings.HasSuffix(cc.Path, ".json") {
		configType = "json"
	} else if strings.HasSuffix(cc.Path, ".ini") {
		configType = "ini"
	} else if strings.HasSuffix(cc.Path, ".properties") {
		configType = "properties"
	} else if strings.HasSuffix(cc.Path, ".yaml") || strings.HasSuffix(cc.Path, ".yml") {
		configType = "yaml"
	} else {
		return
	}
	viper.SetConfigFile(cc.Path)
	viper.SetConfigType(configType)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read ["+cc.Path+"] config is failed", err)
	}
	log.Printf("%v", viper.AllKeys())
}

func ConfigPath() string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		userHome = "."
	}
	workHome := filepath.Join(userHome, ".config-catch/")
	_, err = os.Open(workHome)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("open work home failed!", err)
		}
		os.MkdirAll(workHome, os.ModePerm)
	}
	return filepath.Join(workHome, "cc.json")
}

func ConfigContent() string {
	bytes, err := json.Marshal(cc)
	if err != nil {
		return ""
	}
	return string(bytes)
}
