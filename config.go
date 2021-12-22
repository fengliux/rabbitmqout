package rabbitmqout


import (
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
)

type config struct {
	Codec      codec.Config `config:"codec"`
	Host       string       `config:"host"`
	Port       uint         `config:"port"`
	Username   string       `config:"username"`
	Password   string       `config:"password"`
	Vhost      string       `config:"vhost"`
	Exchange   string       `config:"exchange"`
	RoutingKey string       `config:"routing_key"`
}

var (
	defaultConfig = config{
		Host:       "localhost",
		Port:       5672,
		Username:   "guest",
		Password:   "guest",
		Vhost:      "",
		Exchange:   "filebeat-exc",
		RoutingKey: "filebeat",
	}
)

func (c *config) Validate() error {
	return nil
}
