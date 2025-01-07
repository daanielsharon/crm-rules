package utils

import (
	"bytes"
	"io"
	"net/http"
)

func ForwardRequest(url string, method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	return resp, nil
}
