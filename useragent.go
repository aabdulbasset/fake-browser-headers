package fakeheaders

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func (f *FakeHeaders) RandomUserAgent() (string, error) {
	if len(f.UserAgents) <= 0 {
		return "", errors.New("No User Agents found")
	}
	return f.UserAgents[random(len(f.UserAgents))], nil
}

func (f *FakeHeaders) UpdateAgentsList() {
	downloadUserAgents()
}
func downloadUserAgents() {
	url := "https://raw.githubusercontent.com/EIGHTFINITE/top-user-agents/main/index.json"

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(err)
	}

	// Create or open a local file for writing
	file, err := os.Create("static/useragents.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy the response body to the local file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}

}
