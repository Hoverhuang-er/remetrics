package conf

type Config struct {
	Pushgateway PushGW `yaml,toml,json:"pushgateway"`
}

type PushGW struct {
	GwAddr string `yaml,toml,json:"gw_addr"`
}

type Input struct {
}
