package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func AddRule(rawRule string) (err error) {
	rule := parseRule(rawRule)

	var conf *yamlConfig
	err = yaml.Unmarshal(loadConfig(), &conf)
	if err != nil {
		fmt.Println("yaml unmarshal failed...")
		return 
	}

	conf.Rule = append(conf.Rule, rule)

	y, err := yaml.Marshal(&conf)

	if err != nil {
		fmt.Println("yaml marshal error")
		return
	}

	err = ioutil.WriteFile(ConfigFile, y, 0600)
	if err != nil {
		fmt.Println("write file error")
		return 
	}
	return
}

func parseRule(rawRule string) (r string) {
	rule := strings.Split(rawRule, ",")
	if len(rule) != 3 {
		log.Fatal("invalid rule format")
	}

	// check whether the prefix is supproted and reassign it with the right format.  
	prefix := strings.ToLower(rule[0])
	switch prefix {
		case "domain", "ip-cidr", "geoip":
			prefix = strings.ToUpper(prefix)
		case "keyword", "suffix":
			prefix = "DOMAIN-" + strings.ToUpper(prefix)
		default:
			log.Fatalf("invalid rule prefix: %s", prefix)
	}

	r = strings.Join([]string{prefix, rule[1], rule[2]}, ",")

	return
}


func DeleteRule(domain string) (isDeleted bool) {
	return
}

func ModifyRule(rule string) (isAltered bool) {
	return
}

func SearchDomain(domain string) (isExisted bool) {
	return
}

// check if the given strings can make up a valid proxy rule
func checkFormat(rule string) (isValid bool) {
	return
}

func sortProxy() (isSorted bool) {
	return
}

func loadConfig()[]byte {
	conf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("read config file failed: %s", err)
	} 

	return conf
}