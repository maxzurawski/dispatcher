package config

import (
	"os"

	"github.com/maxzurawski/utilities/config"
	"github.com/maxzurawski/utilities/rabbit"
)

type dispatcherConfigManager struct {
	config.Manager
	proxyService string
	rabbit.RabbitMQManager
}

var instance *dispatcherConfigManager

func DispatcherConfig() *dispatcherConfigManager {
	if instance == nil {
		instance = new(dispatcherConfigManager)
		instance.Init()
	}
	return instance
}

func (c *dispatcherConfigManager) Init() {
	c.Manager.Init()

	if proxyService, err := os.LookupEnv("PROXY_SERVICE"); !err {
		c.proxyService = "http://localhost:8000/api"
	} else {
		c.proxyService = proxyService
	}

	if c.ConnectToRabbit() {
		c.RabbitMQManager.InitConnection(c.RabbitURL())
	}
}

func (c *dispatcherConfigManager) ProxyService() string {
	return c.proxyService
}
