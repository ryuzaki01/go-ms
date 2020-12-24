package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/ryuzaki01/go-ms/encrypt/app/misc"
)

func defaultConfig() Config {
	return Config{
		Name:          "GoMS-Encrypt",
		Port:          3001,
		LogLevel:      4,
		AccessLog:     true,
		AESKey: "9a576e02a3beca3a64f2a1b5fa9061ea7d14716bf205d1b7c256ff7ef511a796",
	}
}

// NewConfig returns a config struct created by
// merging environment variables and a config file.
func NewConfig() *Config {
	temp := environmentConfig()
	config := &temp

	if !config.complete() {
		config.merge(fileConfig())
	}
	defer func() {
		config.merge(defaultConfig())
		config.trimWhitespace()
	}()
	return config
}

func environmentConfig() Config {
	return Config{
		Name:          os.Getenv("APP_NAME"),
		Port:          misc.ParseUint16(os.Getenv("APP_PORT")),
		LogLevel:      misc.Atoi(os.Getenv("APP_LOG_LEVEL")),
		AccessLog:     misc.ParseBool(os.Getenv("APP_ACCESS_LOG")),
		AESKey: 	   os.Getenv("APP_AES_KEY"),
	}
}

func toStringArray(values string) []string {
	if misc.ZeroOrNil(values) {
		return []string{}
	}
	return strings.Split(values, ",")
}

func fileConfig() Config {
	path := misc.NVL(os.Getenv("CONFIG_FILE_PATH"), "/etc/go-ms/config.json")
	file, err := os.Open(path)
	if err != nil {
		return Config{}
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read config file", "err:", err)
		return Config{}
	}
	if strings.TrimSpace(string(data)) == "" {
		return Config{}
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error reading config json data. [message] ", err)
	}
	return config
}

func (config *Config) merge(arg Config) *Config {
	mine := reflect.ValueOf(config).Elem()
	theirs := reflect.ValueOf(&arg).Elem()

	for i := 0; i < mine.NumField(); i++ {
		myField := mine.Field(i)
		if misc.ZeroOrNil(myField.Interface()) {
			myField.Set(reflect.ValueOf(theirs.Field(i).Interface()))
		}
	}
	return config
}

func (config *Config) complete() bool {
	cfg := reflect.ValueOf(config).Elem()

	for i := 0; i < cfg.NumField(); i++ {
		if misc.ZeroOrNil(cfg.Field(i).Interface()) {
			return false
		}
	}
	return true
}

func (config *Config) trimWhitespace() {
	cfg := reflect.ValueOf(config).Elem()
	cfgAttrs := reflect.Indirect(reflect.ValueOf(config)).Type()

	for i := 0; i < cfg.NumField(); i++ {
		field := cfg.Field(i)
		if !field.CanInterface() {
			continue
		}
		attr := cfgAttrs.Field(i).Tag.Get("trim")
		if len(attr) == 0 {
			continue
		}
		if field.Kind() != reflect.String {
			continue
		}
		str := field.Interface().(string)
		field.SetString(strings.TrimSpace(str))
	}
}

// String returns a string representation of the config.
func (config *Config) String() string {
	return fmt.Sprintf(
		"Name: %v, Port: %v, LogLevel: %v, AESKey: %v, AccessLog: %v, ",
		config.Name, config.Port, config.LogLevel, config.AESKey, config.AccessLog)
}