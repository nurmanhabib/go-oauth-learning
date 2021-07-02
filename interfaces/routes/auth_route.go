package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
)

func loginRoute(e *gin.Engine) {
	e.GET("/login", loginFormHandler)
	e.POST("/login", loginHandler)
}

func loginFormHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset: utf-8")
	c.HTML(http.StatusOK, "oauth_login.html", gin.H{})
}

func loginHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionStore.Set("LoggedInUserID", "12345")
	errStore := sessionStore.Save()
	if errStore != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errStore)
		return
	}

	c.Redirect(http.StatusFound, "/oauth/authorize")
}
