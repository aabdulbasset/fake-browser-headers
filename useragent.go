package fakeheaders

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func (f *FakeHeaders) RandomUserAgent() (string, error) {
	choosenBrowser := random(2)
	platform := f.Platforms[random(len(f.Platforms))]
	if choosenBrowser == 0 {
		return GenerateEdge(platform)
	} else if choosenBrowser == 1 {
		return GenerateChrome(platform)
	} else {
		return GenerateFirefox(platform)
	}
}

func GenerateEdge(platform string) (string, error) {
	edgetemplate := "Mozilla/5.0 ({{.Platform}}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{{.ChromeVersion}}.0.0.0 Safari/537.36 Edg/{{.EdgeVersion}}"
	edgeversions := []string{
		"116.0.1938.69",
		"116.0.1938.62",
		"115.0.1901.203",
		"114.0.1823.106",
	}
	edge, err := template.New("edge").Parse(edgetemplate)
	if err != nil {
		return "", err
	}
	edgeVersion := edgeversions[random(len(edgeversions))]
	chromeVersion := strings.Split(edgeVersion, ".")[0]
	var ua bytes.Buffer
	edge.Execute(&ua, map[string]string{"Platform": platform, "EdgeVersion": edgeVersion, "ChromeVersion": chromeVersion})
	return ua.String(), nil
}
func GenerateChrome(platform string) (string, error) {
	chrometemplate := "Mozilla/5.0 ({{.Platform}}) AppleWebKit/537.6 (KHTML, like Gecko) Chrome/{{.ChromeVersion}}.0.0.0 Safari/537.36"
	chromeversions := []string{
		"116",
		"115",
		"114",
		"113",
	}
	chrome, err := template.New("chrome").Parse(chrometemplate)
	if err != nil {
		return "", err
	}
	chromeVersion := chromeversions[random(len(chromeversions))]

	var ua bytes.Buffer
	chrome.Execute(&ua, map[string]string{"Platform": platform, "ChromeVersion": chromeVersion})
	return ua.String(), nil
}
func GenerateFirefox(platform string) (string, error) {
	firefoxtemplate := "Mozilla/5.0 ({{.Platform}}; rv:{{.FirefoxVersion}}.0) Gecko/20100101 Firefox/{{.FirefoxVersion}}.0"
	firefoxversions := []string{
		"117",
		"116",
		"115",
		"114",
	}
	firefox, err := template.New("firefox").Parse(firefoxtemplate)
	if err != nil {
		return "", err
	}
	firefoxVersion := firefoxversions[random(len(firefoxversions))]

	var ua bytes.Buffer
	firefox.Execute(&ua, map[string]string{"Platform": platform, "FirefoxVersion": firefoxVersion})
	return ua.String(), nil
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
