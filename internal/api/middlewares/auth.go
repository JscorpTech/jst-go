package middlewares

import (
	"net/http"
	"strings"

	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/pkg/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(repository interfaces.UserRepositoryPort) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Split(c.Request().Header.Get("Authorization"), " ")
			data, err := utils.ParseJwt(token[1])
			if err != nil {
				return c.JSON(http.StatusForbidden, utils.ErrorResponse(err.Error(), nil))
			}
			user, err := repository.FindByID(data.Sub)
			if err != nil {
				return c.JSON(http.StatusForbidden, utils.ErrorResponse(err.Error(), nil))
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
