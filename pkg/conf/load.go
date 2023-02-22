package conf

import (
	toml "github.com/BurntSushi/toml"
	"log"
	"os"
)

func LoadCfg() *Config {
	var cfg Config
	d3, err := PathExists("config.toml")
	if err != nil {
		log.Printf("load config.toml failed, %v", err)
		return nil
	}
	if err := toml.Unmarshal(d3, &cfg); err != nil {
		log.Printf("unmarshal config.toml failed, %v", err)
		return nil
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
