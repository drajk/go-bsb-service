package http

import (
	"encoding/json"
	"net/http"
)

// AddBodyAndStatusToResponse to prepare the response
func AddBodyAndStatusToResponse(data interface{}, code int, w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)

	if data != nil && bodyAllowedForStatus(code) {
		body, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// copied from http2.go https://github.com/golang/net/blob/master/http2/http2.go
// bodyAllowedForStatus reports whether a given response status code
// permits a body. See RFC 7230, section 3.3.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == 204:
		return false
	case status == 304:
		return false
	}
	return true
}
