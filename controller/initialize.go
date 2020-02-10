package controller

import (
	"fmt"
	"os"
	"path"
)

var Configfile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ := os.Getwd()
		fmt.Println(homeDir)
	}

	Configfile = path.Join(homeDir, ".config/clash/config.yaml")

	//Todo: check if configfile doesn't exists.
}
