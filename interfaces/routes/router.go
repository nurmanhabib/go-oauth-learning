package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/nurmanhabib/go-oauth2/interfaces/middleware"
	"gitlab.com/nurmanhabib/go-oauth2/interfaces/routes/api"
)

type Router struct {
	*middleware.AuthMiddleware
	*middleware.AuthWeb
}

func New() *Router {
	return &Router{
		AuthMiddleware: &middleware.AuthMiddleware{},
		AuthWeb: &middleware.AuthWeb{},
	}
}

func (r *Router) Init() *gin.Engine {
	e := gin.Default()

	// Ping Route
	pingRoute(e)

	// Login Route
	loginRoute(e)

	api.UserRoute(e, r.AuthMiddleware)

	// OAuth 2.0 Server
	oauthServerRoute(e, r.AuthWeb)

	return e
}