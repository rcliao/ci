package github

import "fmt"

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
	return fmt.Sprintf("%s%s?client_id=%s&scope=repo:status&redirect_uri=%s/api/github/callback", baseURL, "/login/oauth/authorize", a.clientID, a.redirectURI)
}

// GetToken uses code from AuthorizationLink to get access token for Github API
func (a *API) GetToken(code string) string {
	return ""
}

func post(path string) {

}
