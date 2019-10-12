package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/xdevices/dispatcher/caches/sensors"
	"github.com/xdevices/dispatcher/config"
	"github.com/xdevices/dispatcher/handlers"
	"github.com/xdevices/dispatcher/publishers"
)

func main() {
	e := echo.New()
	e.GET("/ping", handlers.PingHandler)
	e.POST("/temperature/:uuid/:value", handlers.TemperatureHandler)
	e.Logger.Fatal(e.Start(config.DispatcherConfig().Address()))
}

func init() {
	manager := config.InitDispatcherEurekaManager()
	manager.SendRegistrationOrFail()
	manager.ScheduleHeartBeat(config.DispatcherConfig().ServiceName(), 10)

	publishers.InitLogger()
	_ = sensors.Init(uuid.New().String())
}
