package validators

import (
	"fmt"
	"net/http"
	"strconv"

	e "myapp.com/errors"
)

func ValidateContentType(w http.ResponseWriter, r *http.Request, expected string) {
	v := r.Header.Get("Content-Type")

	if v == "" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(`{"message":"Content-Type not informed"}`))
	}

	if v != expected {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf(`{"message":"Content-Type not supported", "expected": %s}`, expected)))
	}
}

func ValidateContentLength(w http.ResponseWriter, r *http.Request, length int64) {
	size, err := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		w.Write([]byte(e.INTERNAL_SERVER_ERROR))
	}

	if size > length {
		w.Write([]byte(fmt.Sprintf(`{"message":"Content-Length exceded the limit of %dkb"}`, length)))
	}
}
