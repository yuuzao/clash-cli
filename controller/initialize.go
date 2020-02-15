package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	yaml "gopkg.in/yaml.v2"
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

var ConfigFile 	string
var BaseUrl 	string
var Conf 		*yamlConfig


func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ := os.Getwd()
		fmt.Println(homeDir)
	}

	ConfigFile = path.Join(homeDir, ".config/clash/config.yaml")

	//terminate if configfile doesn't exists.
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		log.Fatal("config file does not exist, exit now...")
	}

	unmarshal(&Conf)

	BaseUrl = "http://" + Conf.ExternalController

}

// todo: custom config path
func loadConfig()[]byte {
	conf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("read config file failed: %s", err)
	} 

	return conf
}

func unmarshal(in interface{}){
	err := yaml.Unmarshal(loadConfig(), in)
	if err != nil {
		log.Fatalf("yaml unmarshal error:\n%v", err)
	}
}

func marshal(in interface{}) []byte {
	y, err := yaml.Marshal(in)
	if err != nil {
		log.Fatalf("yaml marshal error:\n%v", err)
	}
	return y
}

func writeToYaml(content []byte) {
	err := ioutil.WriteFile(ConfigFile, content, 0600)
	if err != nil {
		log.Fatalf("write to yaml failed:\n%v", err)
	}
}