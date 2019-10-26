package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/maxzurawski/dispatcher/caches/sensors"
	"github.com/maxzurawski/dispatcher/caches/sensortypes"
	"github.com/maxzurawski/dispatcher/config"
	"github.com/maxzurawski/dispatcher/handlers"
	"github.com/maxzurawski/dispatcher/observers"
	"github.com/maxzurawski/dispatcher/publishers"
)

func main() {

	go observers.ObserveSensorChanges()
	go observers.ObserveSensorTypesChanges()

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
	_ = sensortypes.Init(uuid.New().String())
}
