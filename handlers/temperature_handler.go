package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/xdevices/utilities/stringutils"
)

func TemperatureHandler(c echo.Context) error {

	rawUuid := c.Param("uuid")
	rawValue := c.Param("value")

	isUuidValid := stringutils.IsUuidValid(rawUuid)

	if !isUuidValid {
		msgString := fmt.Sprintf("given uuid: [%s] is not valid uuid", rawUuid)
		log.Warn(msgString)
		return c.JSON(http.StatusBadRequest, errors.New(msgString))
	}

	var value float64
	tempValue, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		msgString := fmt.Sprintf("given value: [%s] is not parsable to float number", rawValue)
		log.Warn(msgString)
		return c.JSON(http.StatusBadRequest, errors.New(msgString))
	} else {
		value = tempValue
	}

	successMsg := fmt.Sprintf("measurement received from sensor with uuid: [%s]. temperature was: [%.2f]", rawUuid, value)
	log.Info(successMsg)
	return c.NoContent(http.StatusAccepted)
}
