package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


func ReloadConfig() {
	data := make(map[string]string)
	data["path"] = ConfigFile

	res := httpReq("PUT", ConfigUrl, data)
	if res.StatusCode != 204 {
		log.Fatalln(res.Status)
	}

}

func ChangeMode(mode string) {
	data := map[string]string {"mode": ""}
	switch strings.ToUpper(mode) {
	case "GLOBAL":
		data["mode"] = "Global"
	case "RULE":
		data["mode"] = "Rule"
	case "DIRECT":
		data["mode"] = "Direct"
	}

	res := httpReq("PATCH", ConfigUrl, data)
	if res.StatusCode != 204 {
		log.Fatalln(res.Status)
	}
}

func SwitchNode(mode string, node string) {
	var path string
	switch mode {
	case "GLOBAL":
		path = "GLOBAL"
	default:
		path = "Proxy"
	}
	data := map[string]string {"name": node}
	url := ProxyUrl + "/" + path
	res := httpReq("PUT", url, data)
	if res.StatusCode != 204 {
		log.Fatalln(res.Status)
	}
}

func ShowStatus() CNode{
	curr := currentNode()
	return curr
}


type CNode struct {
	Mode string
	Node string
}

type Proxy struct {
	All 	[]string 
	Now 	string 
	RType	string
}


func currentNode() CNode {
	reqMode := httpReq("GET", ConfigUrl, nil)
	var modeSettings map[string]interface{}
	var mode string
	readRes(*reqMode, &modeSettings)
	for n, m := range modeSettings {
		if n == "mode"{
			switch t := m.(type) {
			case string:
				mode = t
			default:
				log.Fatalf("unknown mode type: %v", m)
			}
		}
	} 

	if mode == "Direct" {
		return CNode{mode, ""}
	}

	//below is to find out the node marked with field "now".
	reqProxy := httpReq("GET", ProxyUrl, nil)
	var proxySettings map[string]map[string]Proxy
	readRes(*reqProxy, &proxySettings)

	currentNode := make(map[string]CNode)
	for _, proxies := range proxySettings {
		for name, content := range proxies {
			switch name {
			case "GLOBAL":
				currentNode["GLOBAL"] = CNode{"GLOBAL", content.Now}
			case "Proxy":
				currentNode["Proxy"] = CNode{"Rule", content.Now}
			}
		}	
	}

	var c CNode
	switch mode {
	case "Global":
		n := currentNode["Global"].Node
		switch n {
		case "Proxy":
			c.Node = currentNode["Proxy"].Node
		default:
			c.Node = currentNode["Global"].Node	
		}
		c.Mode = "GLOBAL"
	case "Rule":
		c = currentNode["Proxy"]

	}
	return c
}

func readRes(res http.Response, rec interface{}){
	if res.StatusCode != 200 {
		log.Fatalln(res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, rec)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

}

func httpReq(method string, url string, data interface{}) *http.Response {
	var payload io.Reader
	if data != nil {
		p, _ := json.Marshal(data)
		payload = bytes.NewReader(p)

	} else {
		payload = nil
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return res
}