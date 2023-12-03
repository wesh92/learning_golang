package aocget

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Request sends a GET request and returns the response body as a string.
// Also takes a cookie string to authenticate with the AoC server.
func Request(searchURL string, cookieString string) (string, error) {
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

func GetInput(yearString string, dayString string, session string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", yearString, dayString)
	sessionId := fmt.Sprintf("session=%s", session)
	body, err := Request(url, sessionId)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	return string(body), nil
}

func ProvideLines(yearString string, dayString string, session string) ([]string, error) {
	input, err := GetInput(yearString, dayString, session)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	return strings.Split(input, "\n"), nil
}
