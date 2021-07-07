package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"gitlab.com/nurmanhabib/go-oauth2/domain/entities"
	"log"
	"net/http"
)

type AuthWeb struct {}

func (a AuthWeb) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		sessionStore, err := session.Start(r.Context(), w, r)

		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		userID, ok := sessionStore.Get("login_user_id")

		if !ok {
			sessionStore.Set("intended_uri", r.URL.String())
			_ = sessionStore.Save()

			log.Println("intended_uri", r.URL.String())

			c.Redirect(http.StatusFound, "/login")
			return
		}

		// Dummy User
		user := &entities.User{
			ID: int32(userID.(int)),
			Name:     "Habib Nurrahman",
			Username: "nurmanhabib",
			Email:    "nurmanhabib@gmail.com",
			Password: "password123",
		}

		c.Set("user", user)

		c.Next()
	}
}