package global

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type ConfigStruct struct {
	Domains        []string `json:"domains"`
	TrustedProxies []string `json:"trusted_proxies"`
	Cros           struct {
		AllowOrigins     string `json:"allow_origins"`
		AllowCredentials bool   `json:"allow_credentials"`
	}
	DBfile string `json:"dbfile"`
	Mode   string `json:"mode"`
	TLS    bool   `json:"tls"`
}

var Config ConfigStruct

func InitConfig() {
	config_dir := os.Getenv("GOBLOG_CONFIG_PATH")
	config_file := "config.json"
	if config_dir != "" {
		if !strings.HasSuffix(config_dir, "/") {
			config_dir += "/" // Ensure each directory segment ends with a slash
		}
		config_file = config_dir + "config.json"
	}
	data, err := os.ReadFile(config_file)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalln(err)
	}
}
