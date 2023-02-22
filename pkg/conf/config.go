package conf

type Config struct {
	Pushgateway       PushGW `yaml,toml,json:"pushgateway"`
	UseConroutinePool bool   `yaml,toml,json:"use_conroutine_pool"`
}

type PushGW struct {
	GwAddr string `yaml,toml,json:"gw_addr"`
}

type Input struct {
	DataSize []int
}
