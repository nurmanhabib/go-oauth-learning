package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"html/template"
	"net/http"
	"path"
)

func loginRoute(e *gin.Engine) {
	e.GET("/login", loginFormHandler)
	e.POST("/login", loginHandler)
}

func loginFormHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles(path.Join("interfaces", "handler", "oauth_login.html"))

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, gin.H{})

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionStore.Set("login_user_id", 123)
	errStore := sessionStore.Save()
	if errStore != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errStore)
		return
	}

	redirectUri := "/oauth/authorize"
	intendedUri, ok := sessionStore.Get("intended_uri")

	if ok {
		redirectUri = intendedUri.(string)
		sessionStore.Delete("intended_uri")
		_ = sessionStore.Save()
	}

	c.Redirect(http.StatusFound, redirectUri)
}
