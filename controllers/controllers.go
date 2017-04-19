package controllers

import (
	"log"
	"net/http"
	"time"

	"medium/helpers"
	"medium/models"
	"medium/store"
	"medium/templates"
	"medium/urls"
	"medium/utils"

	"github.com/gorilla/schema"
)

func LoginController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.LOGIN
	url_patterns := urls.ReturnURLS()
	if req.Method == "GET" {
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
	if req.Method == "POST" {
		err := req.ParseForm()
		user := new(models.User)
		decoder := schema.NewDecoder()
		err = decoder.Decode(user, req.Form)
		if err != nil {
			utils.CustomTemplateExecute(res, req, controllerTemplate, data)
		} else {
			session, _ := utils.GetValidSession(req)
			session.Values["nickname"] = helpers.StripWhiteSpaces(req.Form["Nickname"][0])
			session.Save(req, res)
			http.Redirect(res, req, url_patterns.HOME_PATH, http.StatusSeeOther)
		}
	}
}

func HomeController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.HOME
	if req.Method == "GET" {
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
}

func AddPostController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.POST_ADD
	url_patterns := urls.ReturnURLS()
	if req.Method == "GET" {
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
	if req.Method == "POST" {
		err := req.ParseForm()
		post := new(models.Post)
		decoder := schema.NewDecoder()
		err = decoder.Decode(post, req.Form)
		log.Println(err)
		if err != nil {
			utils.CustomTemplateExecute(res, req, controllerTemplate, data)
		} else {
			session, _ := utils.GetValidSession(req)
			post.Nickname = session.Values["nickname"].(string)
			post.Time = time.Now()
			err = store.SavePost(post)
			http.Redirect(res, req, url_patterns.POSTS_PATH, http.StatusSeeOther)
		}
	}
}

func ViewPostsController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.POSTS
	if req.Method == "GET" {
		data["posts"], _ = store.GetAllPosts()
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
}

func ViewPostController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.POST_VIEW
	if req.Method == "GET" {
		data["post"], _ = store.GetPost(req.URL.Query().Get(":postid"))
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
}
