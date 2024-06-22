package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func GetInput(day, year int) ([]byte, error) {
	file, err := os.ReadFile("input.txt")
	if err == nil {
		return file, nil
	}

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5fb4fecb996933f2e59abe5cd067a5ff7f7b085a877b1c2eb09f3be29343e0cd71866b239fca250063591ab5c9c334d84e7cd6d3442be1ca13",
	}

	requestURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL %s", requestURL)
	}

	jar.SetCookies(parsedURL, []*http.Cookie{cookie})
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for URL %s", requestURL)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to URL %s", requestURL)
	}

	content, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading response %s", err)
	}

	err = os.WriteFile("input.txt", content, 0777)
	fmt.Println(err)
	return content, nil
}
