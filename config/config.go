package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig   oauth2.Config
	FacebookLoginConfig oauth2.Config
}

var AppConfig Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

const OauthFacebookUrlAPI = "https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token="

func LoadConfig() {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     "958040174312-ro5l4kr6c1b9nd0a65hfhsdnrloq3auf.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8494/google_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	// Oauth configuration for Facebook
	AppConfig.FacebookLoginConfig = oauth2.Config{
		ClientID:     "509113954720967",
		ClientSecret: "ec5edab1ffebe8992863a33b4c60c810",
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "http://localhost:8494/fb_callback",
		Scopes: []string{
			"email",
			"public_profile",
		},
	}
}
