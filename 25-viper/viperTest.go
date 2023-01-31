package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	//viper.SetConfigFile("config/config.yaml")
	viper.SetConfigName(".env")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// For environment variables.
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TEST")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	key := "A1.B"
	fmt.Printf("v[%s]=[%s]\n", key, viper.GetString(key))
	key = "A2.B"
	fmt.Printf("v[%s]=[%s]\n", key, viper.GetString(key))
}
