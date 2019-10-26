package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/maxzurawski/dispatcher/config"
	"github.com/maxzurawski/utilities/net"
)

type IpResponse struct {
	Ip string `json:"ip"`
}

func PingHandler(c echo.Context) error {
	ip, err := net.GetIP(config.DispatcherConfig().IgnoreLoopback())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &IpResponse{Ip: ip})
}
