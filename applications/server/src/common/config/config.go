package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	basePath := getBasePath()
	viper.AddConfigPath(basePath + "/config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("default")
	viper.ReadInConfig()

	env := getEnv()
	if env != "default" {
		mergeConfig(env)
	}
}

func Get(key string) interface{} {

	porp := viper.Get(key)

	if porp == nil {
		panic("Property not found: " + key)
	}

	return porp
}

func UnmarshalKey(key string, rawVal interface{}) {
	err := viper.UnmarshalKey(key, rawVal)

	if err != nil {
		panic(err)
	}
}

func getBasePath() string {
	f, _ := os.Getwd()
	return f
}

func getEnv() string {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		if pair[0] == "ENV" {
			return pair[1]
		}
	}

	return "default"
}

func mergeConfig(env string) {
	viper.SetConfigName(env)
	viper.MergeInConfig()
}
