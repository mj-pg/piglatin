package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
}

func Post(input string) (string, error) {

	reqBody := struct {
		Text string `json:"text"`
	}{
		Text: input,
	}
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("jsonize input: %w", err)
	}
	resp, err := http.Post("http://localhost:3000/piglatins",
		"application/json",
		bytes.NewReader(reqJson))
	if err != nil {
		return "", fmt.Errorf("http POST: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}
	return string(body), nil
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
