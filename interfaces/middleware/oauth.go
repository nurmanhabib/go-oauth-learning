package middleware

import "github.com/gin-gonic/gin"

type OAuth struct {}

func (o OAuth) RedirectDefaultApplication() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}