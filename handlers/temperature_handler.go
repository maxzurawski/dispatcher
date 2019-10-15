package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/xdevices/dispatcher/config"

	"github.com/xdevices/dispatcher/caches/sensortypes"

	"github.com/xdevices/dispatcher/caches/sensors"

	"github.com/google/uuid"
	"github.com/xdevices/dispatcher/publishers"

	"github.com/labstack/echo"
	"github.com/xdevices/utilities/stringutils"
)

func TemperatureHandler(c echo.Context) error {

	rawUuid := c.Param("uuid")
	rawValue := c.Param("value")

	isUuidValid := stringutils.IsUuidValid(rawUuid)

	if !isUuidValid {
		msgString := fmt.Sprintf("given uuid: [%s] is not valid uuid", rawUuid)
		publishers.Logger().Warn(uuid.New().String(), rawUuid, msgString)
		return c.JSON(http.StatusBadRequest, errors.New(msgString))
	}

	// check if sensor is registered
	registeredSensor := sensors.Cache.GetByUuid(rawUuid)
	if registeredSensor == nil {
		msgString := fmt.Sprintf("sensor with uuid: [%s] is not registered", rawUuid)
		publishers.Logger().Error(uuid.New().String(), rawUuid, msgString, "")
		// otherwise return status bad request
		return c.JSON(http.StatusBadRequest, errors.New(msgString))
	}

	var value float64
	tempValue, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		msgString := fmt.Sprintf("given value: [%s] is not parsable to float number", rawValue)
		publishers.Logger().Warn(uuid.New().String(), rawUuid, msgString)
		return c.JSON(http.StatusBadRequest, errors.New(msgString))
	} else {
		value = tempValue
	}

	routingKeyToPublish := sensortypes.TypesCache.GetByType(registeredSensor.Type)
	if routingKeyToPublish == nil {
		msg := fmt.Sprintf("no topic defined for the measurement")
		publishers.Logger().Error(uuid.New().String(), rawUuid, msg, "")
		return c.JSON(http.StatusBadRequest, errors.New(msg))
	}

	publishers.Logger().PublishTemperatureMeasurement(*routingKeyToPublish, config.DispatcherConfig().ServiceName(),
		rawUuid, value)
	return c.NoContent(http.StatusAccepted)
}
