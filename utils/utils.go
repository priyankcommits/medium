package utils

import (
	"net/http"
	"strings"
	"text/template"

	"medium/models"
	"medium/templates"
	"medium/urls"

	"github.com/gorilla/sessions"
)

func CustomTemplateExecute(res http.ResponseWriter, req *http.Request, templateName string, data map[string]interface{}) {
	// Append common templates and data structs and execute template
	tempFunc := template.FuncMap{
		"add_params": func(url string, params ...string) string {
			for _, param := range params {
				// not proud of this
				if strings.Contains(url, "placeholder") {
					url = strings.Replace(url, "placeholder", param, 1)
				} else {
					url = strings.Replace(url, "{"+param+"}", "placeholder", 1)
				}
			}
			return url
		},
	}
	data["url_patterns"] = urls.ReturnURLS()
	session, _ := GetValidSession(req)
	data["nickname"] = session.Values["nickname"].(string)
	t, _ := template.New("").Funcs(tempFunc).ParseFiles(templates.BASE, templateName)
	t.ExecuteTemplate(res, "base.html", data)
}

func AddParamsToUrl(url string, args []models.Kwargs) string {
	// Add params to url using a splice of models.kwargs struct
	for _, arg := range args {
		url = strings.Replace(url, "{"+arg.Key+"}", arg.Value, 1)
	}
	return url
}

func GetValidSession(req *http.Request) (*sessions.Session, error) {
	// Returns a valid authenticated user session
	sessStore := sessions.NewCookieStore([]byte("medium"))
	return sessions.Store.Get(sessStore, req, "medium")
}
