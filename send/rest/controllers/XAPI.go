package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	s "../.."
)

// PublishXAPI :
func PublishXAPI(c echo.Context) error {
	defer func() {
		s.PHE(recover(), s.Cfg.Global.ErrLog, false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body := s.Must(ioutil.ReadAll(c.Request().Body)).([]byte)
	n, termID := s.Xapi(string(body))
	return c.JSON(http.StatusAccepted, s.FSf("%d tuples has been sent, @ %s", n, termID))
}
