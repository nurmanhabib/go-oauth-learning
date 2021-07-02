package routes

import "github.com/gin-gonic/gin"

type Router struct {}

func New() *Router {
	return &Router{}
}

func (r *Router) Init() *gin.Engine {
	e := gin.Default()

	// Ping Route
	pingRoute(e)

	// Login Route
	loginRoute(e)

	// OAuth 2.0 Server
	oauthServerRoute(e)

	return e
}