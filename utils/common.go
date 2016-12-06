package utils

import ()

const (
	configFilePath = "/etc/musubi/config.yml"
	configDirPath  = "/etc/musubi/"
)

type Message struct {
	// Message sender's token.
	Source string

	// Message receiver's token.
	Destination []string

	// Message content.
	Data map[string]string
}
