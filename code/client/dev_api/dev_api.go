package devapi

import "net/http"

type DevAPI struct {
	http  http.Client
	url   string
	token string
}

func NewDevAPI(url string, token string) *DevAPI {
	return &DevAPI{
		http:  http.Client{},
		url:   url,
		token: token,
	}
}
