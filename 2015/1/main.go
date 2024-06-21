package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func main() {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5fb4fecb996933f2e59abe5cd067a5ff7f7b085a877b1c2eb09f3be29343e0cd71866b239fca250063591ab5c9c334d84e7cd6d3442be1ca13",
	}

	requestURL := "https://adventofcode.com/2015/day/1/input"
	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	jar.SetCookies(parsedURL, []*http.Cookie{cookie})
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	var floor int
	var firstBasement int

	for i, rn := range b {
		if rn == '(' {
			floor++
		} else {
			floor--
			if floor < 0 && firstBasement == 0 {
				firstBasement = i + 1
			}
		}
	}

	fmt.Printf("Floor at the end: %d.\nFirst step that enters basement: %d.", floor, firstBasement)
}
