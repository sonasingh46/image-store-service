package decoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Decode the request body to appropriate structure based on content
// type
func DecodeBody(req *http.Request, out interface{}) error {
	cType, err := GetContentType(req)
	if err != nil {
		return err
	}

	if cType != "application/json" {
		return fmt.Errorf("Invalid content type:%s", cType)
	}
	return DecodeJsonBody(req, out)
}

// Get the value of Content-Type that is set in http request header
func GetContentType(req *http.Request) (string, error) {

	if req.Header == nil {
		return "", fmt.Errorf("Request does not have any header")
	}

	return req.Header.Get("Content-Type"), nil
}

// decodeJsonBody is used to decode a JSON request body
func DecodeJsonBody(req *http.Request, out interface{}) error {
	dec := json.NewDecoder(req.Body)
	return dec.Decode(&out)
}
