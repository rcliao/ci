package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiBaseURL = "https://api.github.com"
var baseURL = "https://github.com"

// Event captures the data we want sent by Github webhook
type Event struct {
	Head       commit     `json:"head_commit"`
	Repository repository `json:"repository"`
}

type commit struct {
	ID string `json:"id"`
}

type repository struct {
	Name  string `json:"name"`
	Owner owner  `json:"owner"`
}

type owner struct {
	Name string `json:"name"`
}

type authDTO struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type authResponseDTO struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`

	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// API implements Github related interaction
type API struct {
	clientID     string
	clientSecret string
	redirectURI  string
	token        string
}

// NewAPI is constructor pattern to create API object
func NewAPI(clientID, clientSecret, redirectURI, token string) *API {
	return &API{clientID, clientSecret, redirectURI, token}
}

// UpdateStatus sends the request to create/update status on Github
func (a *API) UpdateStatus(id, state, targetURL, description, context string) error {
	return nil
}

// AuthorizationLink returns link to grant application access token
func (a *API) AuthorizationLink() string {
	return fmt.Sprintf(
		"%s%s?client_id=%s&scope=repo:status&redirect_uri=%s/api/github/callback",
		baseURL,
		"/login/oauth/authorize",
		a.clientID,
		a.redirectURI,
	)
}

// GetToken uses code from AuthorizationLink to get access token for Github API
func (a *API) GetToken(code string) string {
	apiURL := fmt.Sprintf(
		"%s%s",
		baseURL,
		"/login/oauth/access_token",
	)
	body := authDTO{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		Code:         code,
		RedirectURI:  a.redirectURI + "/api/github/callback",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Failed to marshal auth token", err)
		return ""
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Failed to construct request for getting token", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to get response from Github", err)
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read body from Github get token response", err)
		return ""
	}
	var accessResp authResponseDTO
	err = json.Unmarshal(b, &accessResp)
	if err != nil {
		fmt.Println("Failed to parse json body from Github get token response", err)
		return ""
	}
	if accessResp.Error != "" {
		fmt.Println("Failed to retrieve token due to Github error", accessResp)
	}
	return accessResp.AccessToken
}
