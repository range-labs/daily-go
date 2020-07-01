// Package daily exposes methods for interacting with daily.co's REST API.
// https://docs.daily.co/reference
package daily

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "daily-go/" + libraryVersion
	defaultBaseURL = "https://api.daily.co/v1/"
)

// Option defines an option for a client.
type Option func(*Client)

// WithAuth wraps the http client with necessary authentication headers.
func WithAuth(accessToken string) Option {
	return func(c *Client) {
		c.HTTPClient = &authClient{
			httpClient:  c.HTTPClient,
			accessToken: accessToken,
		}
	}
}

// Client for the daily.co API.
type Client struct {
	HTTPClient httpClient
	BaseURL    url.URL
	UserAgent  string
}

// New builds a new Daily client.
func New(opts ...Option) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		HTTPClient: &http.Client{Timeout: time.Second * 5},
		BaseURL:    *baseURL,
		UserAgent:  userAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// GetDomainConfig returns domain configuration information
func (c *Client) GetDomainConfig(ctx context.Context) (*DomainConfig, error) {
	resp := &DomainConfig{}
	return resp, c.request(ctx, "GET", "", nil, resp)
}

// SetDomainConfig updates domain configuration information.
func (c *Client) SetDomainConfig(ctx context.Context, req *Config) (*DomainConfig, error) {
	resp := &DomainConfig{}
	return resp, c.request(ctx, "POST", "", struct {
		Properties *Config
	}{req}, resp)
}

// ListRooms returns available rooms.
func (c *Client) ListRooms(ctx context.Context, req *ListRoomsRequest) (*ListRoomsResponse, error) {
	if req == nil {
		req = &ListRoomsRequest{}
	}
	resp := &ListRoomsResponse{}
	return resp, c.request(ctx, "GET", "rooms", req, resp)
}

// CreateRoom creats a new room.
func (c *Client) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {
	resp := &CreateRoomResponse{}
	return resp, c.request(ctx, "POST", "rooms", req, resp)
}

// GetRoom returns a single room object.
func (c *Client) GetRoom(ctx context.Context, name string) (*GetRoomResponse, error) {
	resp := &GetRoomResponse{}
	return resp, c.request(ctx, "GET", "rooms/"+name, nil, resp)
}

// UpdateRoom updates details about a room.
func (c *Client) UpdateRoom(ctx context.Context, name string, req *UpdateRoomRequest) (*UpdateRoomResponse, error) {
	resp := &UpdateRoomResponse{}
	return resp, c.request(ctx, "POST", "rooms/"+name, req, resp)
}

// DeleteRoom deletes a room.
func (c *Client) DeleteRoom(ctx context.Context, name string) error {
	// Throw away response. It has a 'deleted' property which is always true.
	resp := map[string]interface{}{}
	return c.request(ctx, "DELETE", "rooms/"+name, nil, &resp)
}

func (c *Client) request(ctx context.Context, method, path string, data interface{}, result interface{}) error {
	rel, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("daily: failed to parse request path: %s", err)
	}
	u := c.BaseURL.ResolveReference(rel)

	var body io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("daily: failed to parse request data: %s", err)
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return fmt.Errorf("daily: failed to build request: %s", err)
	}

	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("daily: request failed: %s", err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var msg string
		switch resp.StatusCode {
		case http.StatusBadRequest:
			msg = ErrBadRequest
		case http.StatusUnauthorized:
			msg = ErrUnauthorized
		case http.StatusTooManyRequests:
			msg = ErrTooManyRequests
		case http.StatusInternalServerError:
			msg = ErrInternal
		default:
			msg = ErrUnexpected
		}
		details := &ErrorDetails{}
		if err := json.Unmarshal(respBody, details); err != nil {
			details = nil
		}
		return Error{
			Message:    msg,
			StatusCode: resp.StatusCode,
			Details:    details,
			RawDetails: string(respBody),
		}
	}

	if err = json.Unmarshal(respBody, result); err != nil {
		return Error{
			Message:    ErrParseError + ": " + err.Error(),
			StatusCode: resp.StatusCode,
			RawDetails: string(respBody),
		}
	}

	return nil
}
