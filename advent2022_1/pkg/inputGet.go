package inputGet

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetInput(day int) string {
	url := "https://adventofcode.com/2022/day/" + string(rune(day)) + "/input"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Read the response body
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// Convert the body to a string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	return bodyString
}
