package main

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

func init() {
	runPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(runPath + "/" + "config.yml"); os.IsNotExist(err) {
		panic("config.yml doesn't exists")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(runPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic("config file error")
	}
}

func available(module string) bool {
	config := viper.Get(module)
	if config != nil {
		if reflect.TypeOf(config).String() == "bool" {
			return config.(bool)
		}
		return false
	}
	return false
}

func setPort() string {
	config := viper.Get("port")
	if config != nil {
		if reflect.TypeOf(config).String() == "int" {
			return ":" + strconv.Itoa(config.(int))
		}
		return ":3000"
	}
	return ":3000"
}
