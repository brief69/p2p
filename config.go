package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct {
	P2P struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"p2p"`
}

func loadConfiguration() (*Configuration, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("設定ファイルを開くのに失敗しました: %s", err.Error())
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode config file: %s", err.Error())
	}
	
	
}
