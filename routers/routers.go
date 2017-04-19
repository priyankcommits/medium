package routers

import (
	"net/http"

	"medium/controllers"
	"medium/middlewares"
	"medium/urls"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

// registers all routes for the application.
func GetRouter() *pat.Router {
	// url paths imported from urls package
	url_patterns := urls.ReturnURLS()
	medium := pat.New()
	medium.Get(url_patterns.HOME_PATH, controllers.HomeController)
	medium.Get(url_patterns.POST_ADD_PATH, controllers.AddPostController)
	medium.Post(url_patterns.POST_ADD_PATH, controllers.AddPostController)
	medium.Get(url_patterns.POST_VIEW_PATH, controllers.ViewPostController)
	medium.Get(url_patterns.POSTS_PATH, controllers.ViewPostsController)
	common := pat.New()
	// static route
	common.PathPrefix(url_patterns.STATIC_PATH).Handler(
		http.StripPrefix(url_patterns.STATIC_PATH, http.FileServer(http.Dir("static"))))
	// applying middlewares
	common.PathPrefix(url_patterns.MEDIUM_PATH).Handler(
		negroni.New(
			negroni.HandlerFunc(
				middlewares.LoggingMiddleware),
			negroni.Wrap(medium),
		),
	)
	common.Get(url_patterns.LOGIN_PATH, controllers.LoginController)
	common.Post(url_patterns.LOGIN_PATH, controllers.LoginController)
	return common
}
