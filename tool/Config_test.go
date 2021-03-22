package tool

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestParseConfig(t *testing.T) {
	file, err := os.Open("../config/app.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	cfg := &config{}
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(err)

	}

	log.Fatal(cfg)
}

type config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}
