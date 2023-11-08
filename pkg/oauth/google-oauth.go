package oauth

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupGoogleOauth() *oauth2.Config {
	oauth2GoogleConfig := &oauth2.Config{
		ClientID:     viper.GetString("GOOGLE_OAUTH.CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_OAUTH.CLIENT_SECRET"),
		RedirectURL:  viper.GetString("GOOGLE_OAUTH.CALLBACK_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return oauth2GoogleConfig
}
