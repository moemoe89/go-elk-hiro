//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/practicing-elk-golang/api"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/ping", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = ap.Ping(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
