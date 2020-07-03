package demo

import (
	"demo-service/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (r *Route) Test(c echo.Context) error {
	ctx := &util.CustomEchoContext{Context: c}

	demo, err := r.route.DemoRepo.GetByID(ctx, 1)
	if err != nil {
		return util.Response.Error(ctx, http.StatusInternalServerError, err)
	}

	return util.Response.Success(c, demo)
}
