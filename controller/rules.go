package controller

import (
	// "fmt"
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func AddRule(rawRule string) {
	rule := parseRule(rawRule)

	var conf *yamlConfig
	unmarshal(&conf)
	conf.Rule = append(conf.Rule, rule)

	y := marshal(&conf)
	writeToYaml(y)
}

func parseRule(rawRule string) (r string) {
	/*
	check wether the input rule is valid and return a formatted string.
	"rawRule" is basically made up by three parts: prefix, domain and adapter. Since the 
	prefix is of uppercase and seems a little bit long, this function could handle
	a shorthand or lowercase input of prefix such as "suffix".
	*/

	rule := strings.Split(rawRule, ",")
	if len(rule) != 3 {
		log.Fatal("invalid rule format")
	}

	// check whether the prefix is supported and reassign it with the right format.
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


func DeleteRule(rawRule string) (err error) {
	/*
	Two formats of rawRule are supported: a standard clash rule or just a 
	domain.
	This function will do a search first. If there is only one matched
	rule, it will deleted directly. But if there exist more than one rules
	contain the domain, they will be listed as candidates and wait for the user's
	choice.
	*/
	return
}

func search(rawRule string) (r string) {
	return
}

func ModifyRule(rule string) (isAltered bool) {
	return
}

func SearchDomain(domain string) (isExisted bool) {
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