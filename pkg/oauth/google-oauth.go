package oauth

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
		},
		Endpoint: google.Endpoint,
	}

	return oauth2GoogleConfig
}

func GetResponseAccountGoogle(code string, config *oauth2.Config) (string, error) {

	token, errToken := config.Exchange(context.Background(), code)

	if errToken != nil {
		return "", errors.New("Error exchange google")
	}

	res, ErrRes := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))

	if ErrRes != nil {
		fmt.Println(ErrRes)
		return "", ErrRes
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	fmt.Println(response)

	return string(response), nil

}
