package daily

import "net/http"

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type authClient struct {
	httpClient
	accessToken string
}

func (a *authClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+a.accessToken)
	req.Header.Add("Content-Type", "application/json")
	return a.httpClient.Do(req)
}
