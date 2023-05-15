package client

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
}

func Get() (string, error) {
	resp, err := http.Get("http://localhost:3000/piglatins")
	if err != nil {
		return "", fmt.Errorf("http GET: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}
	return string(body), nil
}
