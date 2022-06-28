package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"gitlab.com/nurmanhabib/go-oauth2/interfaces/middleware"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	authServerURL = "http://localhost:8080"
)

var (
	defaultOauthConfig = oauth2.Config{
		ClientID:     "000000",
		ClientSecret: "999999",
		Scopes:       []string{"view-email", "view-phone", "view-public-profile"},
		RedirectURL:  "https://oauth.pstmn.io/v1/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
)

func oauthServerRoute(e *gin.Engine, aw *middleware.AuthWeb) {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	_ = clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "https://oauth.pstmn.io",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Description)
	})

	o := e.Group("oauth", aw.Authenticate())
	o.GET("/", defaultHandler)
	o.GET("/authorize", func(c *gin.Context) {
		authorizeFormHandler(c, srv)
	})
	o.POST("/approve", func(c *gin.Context) {
		authorizeHandler(c, srv)
	})
	o.POST("/decline", func(c *gin.Context) {

	})
	o.POST("/token", func(c *gin.Context) {
		tokenHandler(c, srv)
	})
}

func defaultHandler(c *gin.Context) {
	u := defaultOauthConfig.AuthCodeURL("ok")
	http.Redirect(c.Writer, c.Request, u, http.StatusFound)
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	userID = "12345"
	err = nil
	return
}

func tokenHandler(c *gin.Context, srv *server.Server) {
	err := srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func authorizeFormHandler(c *gin.Context, srv *server.Server) {
	_, err := srv.ValidationAuthorizeRequest(c.Request)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("interfaces", "handler", "oauth_authorize.html"))

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"name":  "Example App",
		"query": template.URL(c.Request.URL.RawQuery),
	}

	err = tmpl.Execute(c.Writer, data)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func authorizeHandler(c *gin.Context, srv *server.Server) {
	r := c.Request
	w := c.Writer

	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := sessionStore.Get("LoggedInUserID")
	if ok {
		var form url.Values
		if v, ok := sessionStore.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		r.Form = form
	}

	sessionStore.Delete("ReturnUri")
	err = sessionStore.Save()
	if err != nil {
		return
	}

	err = srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	}
}
