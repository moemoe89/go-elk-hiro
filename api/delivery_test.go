//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/go-elk-hiro/api"
	conf "github.com/moemoe89/go-elk-hiro/config"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/error", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := conf.InitLog()
	ctrl := ap.NewCtrl(log)

	err = ctrl.Error(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestInfo(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/info", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := conf.InitLog()
	ctrl := ap.NewCtrl(log)

	err = ctrl.Info(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDebug(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/debug", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := conf.InitLog()
	ctrl := ap.NewCtrl(log)

	err = ctrl.Debug(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestWarn(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/warn", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := conf.InitLog()
	ctrl := ap.NewCtrl(log)

	err = ctrl.Warn(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}