package controller

import (
	"fmt"
	"log"
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

	//terminate if configfile doesn't exists.
	if _, err := os.Stat(Configfile); os.IsNotExist(err) {
		log.Fatal("config file does not exist, exit now...")
	}
}

// todo: custom config path