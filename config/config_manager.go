package config

import "github.com/xdevices/utilities/config"

type dispatcherConfigManager struct {
	config.Manager
}

var instance *dispatcherConfigManager

func DispatcherConfig() *dispatcherConfigManager {
	if instance == nil {
		instance = new(dispatcherConfigManager)
		instance.Init()
	}
	return instance
}
