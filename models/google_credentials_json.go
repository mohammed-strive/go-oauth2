package models

type GoogleCredentials struct {
	Web WebCredentials `json:"web"`
}

type WebCredentials struct {
	ClientID     string   `json:"client_id" required:"true"`
	ClientSecret string   `json:"client_secret" required:"true"`
	ProjectID    string   `json:"project_id" required:"true"`
	RedirectUrls []string `json:"redirect_uris" required:"true"`
}
