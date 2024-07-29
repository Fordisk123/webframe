package conf

import (
	"fmt"
	"github.com/ory/viper"
)

func InitConf(confDir string) {
	if len(confDir) == 0 {
		confDir = "./conf"
	}
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(confDir)  // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
