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
		Value: "53616c7465645f5fe6c2a746dd8425143171d518691709286b995f82d5c6ba84a77e8cf5a8e6ec6752fcf910593724cdd7f108773de3d5136d8c01764bbc00d0",
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

func Contains[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}

	return false
}
