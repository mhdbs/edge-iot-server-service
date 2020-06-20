package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	HTTP struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"http"`
}

var configFile = "/var/edge-iot-server-service.json"
var HttpHost string
var HttpPort string

func LoadConfig() {
	readJSON, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(">>>", err)
	}
	configData := Config{}
	err = json.Unmarshal([]byte(readJSON), &configData)
	if err != nil {
		fmt.Println("error on json unmarshall", err)
	}
	fmt.Println(configData)
	HttpHost = configData.HTTP.Host
	HttpPort = configData.HTTP.Port
}
