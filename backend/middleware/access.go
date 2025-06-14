package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type AccessMiddlewareRule struct {
	Role       *string
	Permission *string
	AllowSelf  bool
	SelfIDKey  *string
}

func (r *AccessMiddlewareRule) CheckRule(c echo.Context) (bool, error) {
	if r.Role != nil {
		roles := c.Get("roles").([]string)
		for _, role := range roles {
			if role == *r.Role {
				return true, nil
			}
		}
		return false, errors.New("unauthorized, need role " + *r.Role)
	}
	if r.Permission != nil {
		// TODO: implement
	}
	if r.AllowSelf && r.SelfIDKey != nil {
		requesterID := c.Get("userID").(uint)
		userID, ok := c.Get(*r.SelfIDKey).(uint)
		if !ok {
			return false, errors.New("unauthorized, need self")
		}
		if requesterID == userID {
			return true, nil
		}
		return false, errors.New("unauthorized, need self")
	}
	return false, nil
}

func AccessMiddleware(rules []AccessMiddlewareRule) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, rule := range rules {
				ok, err := rule.CheckRule(c)
				if err != nil {
					return c.JSON(401, err.Error())
				}
				if ok {
					return next(c)
				}
			}
			return next(c)
		}
	}
}
