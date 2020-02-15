package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"bytes"
)

func ReloadConfig() {
	configUrl := BaseUrl + "/configs"

	data := make(map[string]string)
	data["path"] = ConfigFile
	payload, _ := json.Marshal(data)


	client := http.Client{}
	req, err := http.NewRequest("PUT", configUrl, bytes.NewReader((payload)))
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
}
