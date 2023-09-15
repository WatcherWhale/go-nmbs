package irail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const IRAIL_API string = "https://api.irail.be/"

type Request[Result any] struct {
	Path       string
	Parameters map[string]string
}

type ErrorObject struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func (r *Request[Result]) parseParameter() string {
	params := url.Values{}

	for key, value := range r.Parameters {
		params.Add(key, value)
	}

	return params.Encode()
}

func (r *Request[Result]) Do(resultObj *Result) error {
	r.Parameters["format"] = "json"

	resp, err := http.Get(IRAIL_API + r.Path + "/?" + r.parseParameter())

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	errorObj := ErrorObject{
		Error: -1,
	}
	err = json.Unmarshal(body, &errorObj)

	if err != nil {
		return err
	}

	if errorObj.Error != -1 {
		return fmt.Errorf(errorObj.Message)
	}

	err = json.Unmarshal(body, resultObj)

	if err != nil {
		return err
	}

	return nil
}
