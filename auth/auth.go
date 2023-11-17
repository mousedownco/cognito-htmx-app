package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type Cognito struct {
	Endpoint    string
	ClientId    string
	RedirectUri string
}

func NewCognito(endpoint, clientId, redirectUri string) *Cognito {
	return &Cognito{
		Endpoint:    endpoint,
		ClientId:    clientId,
		RedirectUri: redirectUri,
	}
}

func (c *Cognito) CodeForToken(code string) (*TokenResponse, error) {
	req := url.Values{}
	req.Set("grant_type", "authorization_code")
	req.Set("client_id", c.ClientId)
	req.Set("code", code)
	req.Set("redirect_uri", c.RedirectUri)

	resp, e := http.PostForm(c.Endpoint, req)
	if e != nil {
		return nil, e
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var token TokenResponse
	e = json.NewDecoder(resp.Body).Decode(&token)
	return &token, e
}

func HandleCognitoCallback(c *Cognito, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
			return
		}
		token, e := c.CodeForToken(code)
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("token: %+v", token)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}
