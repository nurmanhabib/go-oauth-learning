package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware struct {}

type BearerToken struct {
	Raw string `header:"Authorization"`
}

func (b BearerToken) Token() string {
	return strings.TrimSpace(b.Raw[len("Bearer"):])
}

func (a AuthMiddleware) RedirectIfAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token BearerToken

		err := c.ShouldBindHeader(&token)

		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusUnauthorized)
			return
		}

		c.Set("token", token)

		log.Print(token.Token())

		c.Next()
	}
}
