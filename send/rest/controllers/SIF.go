package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	s "../.."
)

// PublishSIF :
func PublishSIF(c echo.Context) error {
	defer func() {
		s.PHE(recover(), s.Cfg.Global.ErrLog, false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body := s.Must(ioutil.ReadAll(c.Request().Body)).([]byte)
	nV, nS, nA, termID := s.Sif(string(body))
	return c.JSON(http.StatusAccepted, s.FSf("<%d> value tuples, <%d> struct tuples, <%d> array info tuples have been sent, @ %s", nV, nS, nA, termID))
}
