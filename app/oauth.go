package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type callbackResponse struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Scope         string `json:"scope"`
	Expires_in    int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
}

func (a *App) Callback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		code := r.URL.Query().Get("code")
		if code == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Bad request")
			return
		}

		accessTokenResponse, err := requestAccessToken(code, a.conf.ClientID, a.conf.ClientSecret)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(accessTokenResponse)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accessTokenResponse)
	}
}

func requestAccessToken(code string, clientID string, clientSecret string) (callbackResponse, error) {
	var response callbackResponse

	url := "https://accounts.spotify.com/api/token"
	redirect_uri := "http%3A%2F%2F127.0.0.1%3A8000%2Fredirect-URI"
	payload := strings.NewReader("grant_type=authorization_code&" + "code=" + code + "&redirect_uri=" + redirect_uri)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		return response, errors.New("couldn't get access token")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &response)

	fmt.Println("token", response.Access_token)

	return response, nil

}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (a *App) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := generateRandomString(16)

		u := fmt.Sprintf("https://accounts.spotify.com/authorize?%s", url.Values{
			"response_type": {"code"},
			"client_id":     {a.conf.ClientID},
			"redirect_uri":  {a.conf.RedirectURI},
			"state":         {state},
		}.Encode())
		log.Println(u)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}
