// package doppler provices the bindings for the Doppler REST APIs.
package doppler

import "net/http"

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
