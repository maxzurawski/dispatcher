package main

import (
	"github.com/labstack/echo"
	"github.com/xdevices/dispatcher/config"
	"github.com/xdevices/dispatcher/handlers"
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
}
