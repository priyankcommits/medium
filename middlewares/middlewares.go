package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func LoggingMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	// Create return string
	var request []string
	// Time
	request = append(request, "\n")
	timeStamp := time.Now()
	request = append(request, timeStamp.String())
	// Add the request string
	url := fmt.Sprintf("%v %v %v", req.Method, req.URL, req.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", req.Host))
	// Loop through headers
	for name, headers := range req.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if req.Method == "POST" {
		req.ParseForm()
		request = append(request, "\n")
		request = append(request, req.Form.Encode())
	}
	// Log the request as a string

	_, _ = f.WriteString(strings.Join(request, "\n"))
	next(res, req)
}
