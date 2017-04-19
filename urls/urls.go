package urls

type (
	urls struct {
		STATIC_PATH    string
		LOGIN_PATH     string
		HOME_PATH      string
		MEDIUM_PATH    string
		POST_ADD_PATH  string
		POST_VIEW_PATH string
		POSTS_PATH     string
	}
)

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.LOGIN_PATH = "/"
	url_patterns.MEDIUM_PATH = "/medium/"
	url_patterns.HOME_PATH = url_patterns.MEDIUM_PATH + "home/"
	url_patterns.POST_ADD_PATH = url_patterns.MEDIUM_PATH + "form/"
	url_patterns.POSTS_PATH = url_patterns.MEDIUM_PATH + "posts/"
	url_patterns.POST_VIEW_PATH = url_patterns.MEDIUM_PATH + "post/{postid}/"
	return url_patterns
}
