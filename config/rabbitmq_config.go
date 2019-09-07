package config

import "github.com/anypick/infra/base/props/container"

const (
	DefaultPrefix = "rabbit"
)

func init() {
	container.RegisterYamContainer(&RabbitMQConfig{Prefix: DefaultPrefix})
}

type RabbitMQConfig struct {
	Prefix   string
	IpAddr   string `yaml:"ipAddr"`   // 主机ip
	Port     int    `yaml:"port"`     // 端口
	Vhost    string `yaml:"vhost"`    // vhost
	UserName string `yaml:"username"` // string
	Password string `yaml:"password"` // string
}

func (r RabbitMQConfig) ConfigAdd(config map[interface{}]interface{}) {
	r.IpAddr = config["ipAddr"].(string)
	r.Port = config["port"].(int)
	r.Vhost = config["vhost"].(string)
	r.UserName = config["username"].(string)
	r.Password = config["password"].(string)
}
