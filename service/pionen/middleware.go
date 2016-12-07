package pionen

import (
	"github.com/Hackform/hfse/model/pionenmodel"
	"github.com/Hackform/hfse/service/pionen/access"
	"github.com/labstack/echo"
	"net/http"
)

func (p *Pionen) MiddleAuthLevel(level uint8) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := new(pionenmodel.RequestJWT)
			c.Bind(req)
			if p.VerifyJWT(req.Value.Token, level) {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
		}
	}
}

func (p *Pionen) MiddleAuthRoot() echo.MiddlewareFunc {
	return p.MiddleAuthLevel(access.ROOT)
}

func (p *Pionen) MiddleAuthAdmin() echo.MiddlewareFunc {
	return p.MiddleAuthLevel(access.ADMIN)
}

func (p *Pionen) MiddleAuthMod() echo.MiddlewareFunc {
	return p.MiddleAuthLevel(access.MOD)
}

func (p *Pionen) MiddleAuthUser() echo.MiddlewareFunc {
	return p.MiddleAuthLevel(access.USER)
}
