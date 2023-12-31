package router

import (
	"encoding/gob"
	"go-auth0-web/platform/authenticator"
	"go-auth0-web/platform/middleware"
	"go-auth0-web/web/app/callback"
	"go-auth0-web/web/app/login"
	"go-auth0-web/web/app/logout"
	"go-auth0-web/web/app/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// register custom type used by cookies
	// a map literal, with key values and interface values
	// interface values here mean they can be any type
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	/* Continue at "Logging in" */

	return router
}
