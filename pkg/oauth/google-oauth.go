package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"woman-center-be/internal/web/resources/v1"

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
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return oauth2GoogleConfig
}

func GetResponseAccountGoogle(code string, config *oauth2.Config) (*resources.UserGoogleInfo, error) {

	token, errToken := config.Exchange(context.Background(), code)

	if errToken != nil {
		return nil, errors.New("Error exchange google")
	}

	res, ErrRes := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))

	if ErrRes != nil {
		fmt.Println(ErrRes)
		return nil, ErrRes
	}

	defer res.Body.Close()

	var User *resources.UserGoogleInfo

	err := json.NewDecoder(res.Body).Decode(&User)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return User, nil

}
