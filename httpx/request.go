package httpx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// decode json request body with validation
func DecodeJSON(write http.ResponseWriter, read *http.Request, dst any, maxBytes int64) error {
	// enforce a maximum body size
	read.Body = http.MaxBytesReader(write, read.Body, maxBytes)

	// json decoder
	dec := json.NewDecoder(read.Body)
	dec.DisallowUnknownFields() // optional, stricter validation

	if err := dec.Decode(dst); err != nil {
		// normalize error types
		var syntaxErr *json.SyntaxError
		var unmarshalErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalErr):
			return errors.New("body contains invalid value for a field")
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		default:
			return errors.New("invalid JSON body")
		}
	}

	// check for trailing data
	if dec.More() {
		return errors.New("body must contain only a single JSON object")
	}

	return nil
}
