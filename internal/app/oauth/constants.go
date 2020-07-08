package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var (
	USER_SERVICE_URL  = "http://192.168.113.23:8888/api/v1/users"
	TOKEN_LOGIN_URL   = "http://localhost:8888/#/token-login"
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8888/api/v1/oauth/google/callback",
		ClientID:     "196164425223-jg0g4u8h9ve6qukq064kood9ofufstov.apps.googleusercontent.com", // os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
		ClientSecret: "vs0Z7DhShnO6c1o8yiuj4MXo",                                                 // os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)
