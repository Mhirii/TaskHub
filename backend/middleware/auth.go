package middleware

import (
	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"github.com/labstack/echo/v4"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h := c.Request().Header.Get("Authorization")
		if h == "" {
			return c.JSON(401, "unauthorized")
		}
		if h[0:7] != "Bearer " {
			return c.JSON(401, "invalid token")
		}
		token := h[7:]
		t, err := jwt.Parse([]byte(token),
			jwt.WithKey(jwa.HS512(), []byte(config.GetConfig().Auth.Secret)),
		)
		if err != nil {
			return c.JSON(401, "invalid token, invalid signature")
		}
		userID, ok := t.Subject()
		if !ok {
			return c.JSON(401, "invalid token, missing subject")
		}

		var roles []string
		err = t.Get("roles", roles)
		if err != nil {
			return c.JSON(401, "invalid token, missing roles")
		}
		c.Set("userID", userID)
		c.Set("roles", roles)
		c.Set("token", h[7:])
		return next(c)
	}
}
