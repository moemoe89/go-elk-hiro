//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/practicing-elk-golang/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry) *echo.Echo {
	r := echo.New()
	r.Use(middleware.Recover())

	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	ctrl := ap.NewCtrl(log)

	r.GET("/error", ctrl.Error)
	r.GET("/info", ctrl.Info)
	r.GET("/debug", ctrl.Debug)
	r.GET("/warn", ctrl.Warn)

	return r
}
