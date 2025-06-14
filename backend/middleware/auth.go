package middleware

import (
	"fmt"

	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"github.com/Mhirii/TaskHub/backend/pkg/roles"
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
		if len(h) < 12 {
			return c.JSON(401, "invalid token")
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

		var r interface{}
		err = t.Get("roles", &r)
		if err != nil {
			fmt.Println(err)
			return c.JSON(401, "invalid token, missing roles")
		}

		rstr := fmt.Sprintf("%v", r)
		fmt.Println("roles: ", rstr)

		c.Set("userID", userID)
		c.Set("roles", roles.StringToRoles(rstr))
		c.Set("token", h[7:])
		return next(c)
	}
}
