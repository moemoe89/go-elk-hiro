//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(c echo.Context) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	return c.JSON(http.StatusOK, res)
}
