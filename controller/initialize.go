package controller

import (
	"fmt"
	"log"
	"os"
	"path"

)

type yamlConfig struct {
	Port               int          			`yaml:"port"`
	SocksPort          int          			`yaml:"socks-port"`
	RedirPort          int          			`yaml:"redir-port"`
	Authentication     []string     			`yaml:"authentication"`
	AllowLan           bool         			`yaml:"allow-lan"`
	BindAddress        string       			`yaml:"bind-address"`
	Mode               string				    `yaml:"mode"`
	LogLevel           string				 	`yaml:"log-level"`
	ExternalController string       			`yaml:"external-controller"`
	ExternalUI         string       			`yaml:"external-ui"`
	Secret             string       			`yaml:"secret"`

	ProxyProvider map[string]map[string]interface{} `yaml:"proxy-provider"`
	Hosts         map[string]string                 `yaml:"hosts"`
	DNS           map[string]interface{}            `yaml:"dns"`
	Experimental  map[string]interface{}            `yaml:"experimental"`
	Proxy         []map[string]interface{}          `yaml:"Proxy"`
	ProxyGroup    []map[string]interface{}          `yaml:"Proxy Group"`
	Rule          []string                          `yaml:"Rule"`
}

var ConfigFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ := os.Getwd()
		fmt.Println(homeDir)
	}

	ConfigFile = path.Join(homeDir, ".config/clash/configb.yaml")

	//terminate if configfile doesn't exists.
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		log.Fatal("config file does not exist, exit now...")
	}
}

// todo: custom config path