package middlewares

import (
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(repository domain.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Split(c.Request().Header.Get("Authorization"), " ")
			data, err := utils.ParseJwt(token[1])
			if err != nil {
				return c.JSON(http.StatusForbidden, domain.ErrorResponse(err.Error(), nil))
			}
			user, err := repository.FindById(data.Sub)
			if err != nil {
				return c.JSON(http.StatusForbidden, domain.ErrorResponse(err.Error(), nil))
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
