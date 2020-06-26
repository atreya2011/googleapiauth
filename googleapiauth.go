package googleapiauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Retrieves a token, saves the token, then returns the token source to create new service.
// If modifying these scopes, delete your previously saved token.json.
func GetTokenSource(ctx context.Context, credFilename string, scope ...string) (tokenSource oauth2.TokenSource, err error) {
	b, err := ioutil.ReadFile(credFilename)
	if err != nil {
		err = fmt.Errorf("unable to read client secret file: %v", err)
		return
	}

	c, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		err = fmt.Errorf("unable to parse client secret file to config: %v", err)
		return
	}

	token, err := tokenFromFile()
	if err != nil {
		token, err = getTokenFromWeb(c)
		if err != nil {
			return
		}
		if err = saveToken(token); err != nil {
			return
		}
	}

	tokenSource = c.TokenSource(ctx, token)
	return
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (token *oauth2.Token, err error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err = fmt.Scan(&authCode); err != nil {
		err = fmt.Errorf("Unable to read authorization code: %v", err)
		return
	}

	token, err = config.Exchange(context.TODO(), authCode)
	if err != nil {
		err = fmt.Errorf("Unable to retrieve token from web: %v", err)
		return
	}
	return
}

// Retrieves the token from locally saved token.json.
func tokenFromFile() (token *oauth2.Token, err error) {
	f, err := os.Open("token.json")
	if err != nil {
		return
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&token)
	return
}

// Saves the token to token.json file.
func saveToken(token *oauth2.Token) (err error) {
	fmt.Printf("Saving OAuth token to token.json")
	f, err := os.OpenFile("token.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	if err = json.NewEncoder(f).Encode(token); err != nil {
		return
	}
	return
}
