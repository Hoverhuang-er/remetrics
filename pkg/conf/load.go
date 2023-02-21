package conf

import (
	"encoding/json"
	toml "github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadCfg() *Config {
	var cfg Config
	if d, err := PathExists("config.yaml"); err == nil {
		if err := yaml.Unmarshal(d, &cfg); err != nil {
			log.Fatal(err)
			return nil
		}
	} else if d2, err := PathExists("config.json"); err == nil {
		if err := json.Unmarshal(d2, &cfg); err != nil {
			log.Fatal(err)
			return nil
		}
	} else if d3, err := PathExists("config.toml"); err == nil {
		if err := toml.Unmarshal(d3, &cfg); err != nil {
			log.Fatal(err)
			return nil
		}
	}
	return &cfg
}

func PathExists(path string) ([]byte, error) {
	fi, err := os.Stat(path)
	if err == nil {
		log.Printf("file exists, %v", fi)
		return os.ReadFile(path)
	} else if os.IsNotExist(err) {
		log.Print("file not exists")
		return nil, err
	} else {
		log.Print("file not exists")
		return nil, err
	}
}
