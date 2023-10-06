package aocGet

import (
	"fmt"
	"io"
	"net/http"
)

// Request sends a GET request and returns the response body as a string.
func AoCRequest(searchURL string, cookieString string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", err
	}

	// Headers
	req.Header.Set("Cookie", cookieString)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("authority", "adventofcode.com")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response status is successful.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response status: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
