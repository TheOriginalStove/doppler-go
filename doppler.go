// package doppler provices the bindings for the Doppler REST APIs.
package dopplergo

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	// APIVersion is the supported API version.
	APIVersion string = "v3"

	APIURL = "https://api.doppler.com"
)

type APIResponse struct {
	// Contains a map of all HTTP header keys to values.
	Header http.Header

	//RawJSON contains the response body as raw bytes.
	RawJSON []byte

	// StatusCode is the status codea as an integer. (200, 404, 501, etc)
	StatusCode int
}

type Doppler struct {
	Project     string
	Environment string
	Config      string
}

// ClientConfig is created to configure the client in order to access Doppler.
type ClientConfig struct {
	lock sync.RWMutex

	// The HTTP client to use.
	HttpClient *http.Client

	// MinRetryWait is used to control the minimum time to wait before retrying when
	// a 5XX error occurs. Defaults to 1000 milliseconds
	MinRetryWait time.Duration

	// MaxRetryWait controls the maximum time to wait before retrying after a 5XX error.
	// Defaults to 1500 milliseconds.
	MaxRetryWait time.Duration

	// MaxRetries controls the maximum number of times to retry when a 5xx error occurs.
	// To disable retries, set to 0
	MaxRetries int

	// Timeout is for setting a custom timeout in the HttpClient
	Timeout time.Duration

	// If there is something misconfigured, this will be the error
	Error error

	// Limiter is the rate limiter used by the client. If this is nil, no limit will be set.
	Limiter *rate.Limiter
}

func DefaultConfig() *Config {
	return &Config{}
}

type wrappingFuncHandler func(operation, path string) string

type Workplace struct {
	Name         string `json:"name"`
	BillingEmail string `json:"billing_email"`
}

type ActivityLog struct {
	Id          string
	Text        string
	Html        string
	User        User
	Project     string
	Environment string
	Config      string
	CreatedAt   time.Time
	Diff        Diff
}

type User struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	ProfileImgUrl string `json:"profile_image_url"`
}

type Diff struct {
	Name    string
	Added   []string
	Removed []string
	Updated []string
}

////////////////////////////////////////////////////////////////////////////////

// Client is the client to the Doppler API. Create with NewClient
type Client struct {
	lock                sync.RWMutex
	address             *url.URL
	config              *Config
	token               string
	headers             http.Header
	wrappingFuncHandler wrappingFuncHandler
}

// GetActivityLogs will take
func (c *Client) GetActivityLogs() ([]ActivityLog, error) {
	return nil, ErrNotImplemented
}

func (c *Client) GetActivityLog(id string) (ActivityLog, error) {
	return ActivityLog{}, ErrNotImplemented
}

func (c *Client) GetWorkplace() (Workplace, error) {
	return Workplace{}, ErrNotImplemented
}
