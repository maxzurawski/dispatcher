package observers

import (
	"encoding/json"
	"fmt"

	"github.com/maxzurawski/dispatcher/caches/sensortypes"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/dispatcher/config"
	"github.com/maxzurawski/dispatcher/publishers"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/domain"
)

func ObserveSensorTypesChanges() {
	observer := config.DispatcherConfig().InitObserver()
	defer observer.Channel.Close()
	observer.DeclareTopicExchange(crosscutting.TopicConfigurationChanged.String())
	observer.BindQueue(observer.Queue, crosscutting.RoutingKeySensorTypes.String()+".#", crosscutting.TopicConfigurationChanged.String())
	deliveries := observer.Observe()

	for msg := range deliveries {
		confMsg := domain.ConfigurationChanged{}
		err := json.Unmarshal(msg.Body, &confMsg)
		if err != nil {
			publishers.Logger().Error(uuid.New().String(), "", "could not update sensor types cache", err.Error())
			continue
		}
		log.Info(fmt.Sprintf("Received: [%s]\n", string(msg.Body)))
		log.Info(fmt.Sprintf("Routing key: [%s]", msg.RoutingKey))
		err = sensortypes.Init(confMsg.ProcessId)
		if err == nil {
			publishers.Logger().Info(confMsg.ProcessId, "", "successfully updated sensor types cache")
		}
	}
}
