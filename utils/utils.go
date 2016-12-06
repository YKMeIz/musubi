package utils

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/nrechn/musubi/database"
	"github.com/spf13/viper"
)

// LoadConfig loads the config file of this program.
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(configDirPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetString returns value of the key in config file.
func GetString(arg string) string {
	return viper.GetString(arg)
}

// GetStringSlice returns value of the key in config file.
func GetStringSlice(arg string) []string {
	return viper.GetStringSlice(arg)
}

// CheckConfig detects if config file is exist.
func CheckConfig() error {
	if _, err := os.Stat(configFilePath); err != nil {
		return err
	}
	return nil
}

// Secure checks if user would like to force disable TLS/SSL feature.
func Secure() bool {
	s := GetString("secure")
	if s == "disabled" || s == "disable" {
		return false
	}
	return true
}

// ToDo:
func Crypto() bool {
	s := GetString("crypto")
	if s == "disabled" || s == "disable" {
		return false
	}
	return true
}

// Websocket checks if user would like to use alternative push method.
func Websocket() bool {
	s := GetString("pushMethod")
	if s == "websocket" || s == "" {
		return true
	}
	return false
}

// IsSet checks if a givin key has been set in config file.
func IsSet(arg string) bool {
	return viper.IsSet(arg)
}

// ReadJson reads a string of Message content, and returns in Message type.
func ReadJson(msg string) Message {
	var m Message
	dec := json.NewDecoder(strings.NewReader(msg))
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
	}
	return m
}

// Authenticate checks if device's name and token are same
// as record in database.
func Authenticate(name, token string) bool {
	if !database.IsName(name) {
		return false
	}
	if !database.CompareToken(name, token) {
		return false
	}
	return true
}
