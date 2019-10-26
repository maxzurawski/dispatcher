package observers

import (
	"encoding/json"
	"fmt"

	sensors2 "github.com/xdevices/dispatcher/caches/sensors"

	"github.com/xdevices/dispatcher/publishers"

	"github.com/xdevices/dispatcher/config"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/domain"
)

func ObserveSensorChanges() {
	observer := config.DispatcherConfig().InitObserver()
	defer observer.Channel.Close()
	observer.DeclareTopicExchange(crosscutting.TopicConfigurationChanged.String())
	observer.BindQueue(observer.Queue, crosscutting.RoutingKeySensors.String()+".#", crosscutting.TopicConfigurationChanged.String())
	deliveries := observer.Observe()

	for msg := range deliveries {
		confMsg := domain.ConfigurationChanged{}
		err := json.Unmarshal(msg.Body, &confMsg)
		if err != nil {
			publishers.Logger().Error(uuid.New().String(), "", "could not update sensors cache", err.Error())
			continue
		}
		log.Info(fmt.Sprintf("Received: [%s]\n", string(msg.Body)))
		log.Info(fmt.Sprintf("Routing key: [%s]", msg.RoutingKey))
		err = sensors2.Init(confMsg.ProcessId)
		if err == nil {
			publishers.Logger().Info(confMsg.ProcessId, "", "successfully updated sensors cache")
		}
	}
}
