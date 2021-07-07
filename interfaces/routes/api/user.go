package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/nurmanhabib/go-oauth2/interfaces/middleware"
)

func UserRoute(e *gin.Engine, am *middleware.AuthMiddleware) {
	e.GET("api/user", am.RedirectIfAuthenticated(), func(c *gin.Context) {
		var token middleware.BearerToken

		if v, ok := c.Get("token"); ok {
			token = v.(middleware.BearerToken)
		}

		_, _ = c.Writer.Write([]byte(token.Token()))
	})
}
