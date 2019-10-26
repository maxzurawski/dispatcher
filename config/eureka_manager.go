package config

import "github.com/maxzurawski/utilities/discovery"

type DispatcherEurekaManager struct {
	discovery.Manager
}

func InitDispatcherEurekaManager() *DispatcherEurekaManager {
	manager := DispatcherEurekaManager{
		Manager: discovery.Manager{
			RegistrationTicket: DispatcherConfig().RegistrationTicket(),
			EurekaService:      DispatcherConfig().EurekaService(),
		},
	}
	return &manager
}
