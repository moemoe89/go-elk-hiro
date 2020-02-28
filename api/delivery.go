//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry) *ctrl {
	return &ctrl{log}
}

// Error will send the error log
func (ct *ctrl) Error(c echo.Context) error {
	ct.log.Errorf("Error log")
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

// Info will send the info log
func (ct *ctrl) Info(c echo.Context) error {
	ct.log.Infof("Info log")
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

// Debug will send the debug log
func (ct *ctrl) Debug(c echo.Context) error {
	ct.log.Debugf("Debug log")
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

// Warn will send the warn log
func (ct *ctrl) Warn(c echo.Context) error {
	ct.log.Warnf("Warn log")
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}