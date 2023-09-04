package fakeheaders

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func (f *FakeHeaders) RandomUserAgent() (string, error) {

	platform := f.Platforms[random(len(f.Platforms))]
	if f.Browser == Edge {
		return f.GenerateEdge(platform)
	} else if f.Browser == Chrome {
		return f.GenerateChrome(platform)
	} else if f.Browser == Firefox {
		return f.GenerateFirefox(platform)
	} else {
		return "", errors.New("No Browser found, available types are 'chrome', 'firefox', 'edge'")
	}
}

func (f FakeHeaders) GenerateEdge(platform string) (string, error) {
	edgetemplate := "Mozilla/5.0 ({{.Platform}}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{{.ChromeVersion}}.0.0.0 Safari/537.36 Edg/{{.EdgeVersion}}"

	edge, err := template.New("edge").Parse(edgetemplate)
	if err != nil {
		return "", err
	}
	edgeVersion := f.EdgeVersions[random(len(f.EdgeVersions))]
	chromeVersion := strings.Split(edgeVersion, ".")[0]
	var ua bytes.Buffer
	edge.Execute(&ua, map[string]string{"Platform": platform, "EdgeVersion": edgeVersion, "ChromeVersion": chromeVersion})
	return ua.String(), nil
}
func (f FakeHeaders) GenerateChrome(platform string) (string, error) {
	chrometemplate := "Mozilla/5.0 ({{.Platform}}) AppleWebKit/537.6 (KHTML, like Gecko) Chrome/{{.ChromeVersion}}.0.0.0 Safari/537.36"

	chrome, err := template.New("chrome").Parse(chrometemplate)
	if err != nil {
		return "", err
	}
	chromeVersion := f.ChromeVersions[random(len(f.ChromeVersions))]

	var ua bytes.Buffer
	chrome.Execute(&ua, map[string]string{"Platform": platform, "ChromeVersion": chromeVersion})
	return ua.String(), nil
}
func (f FakeHeaders) GenerateFirefox(platform string) (string, error) {
	firefoxtemplate := "Mozilla/5.0 ({{.Platform}}; rv:{{.FirefoxVersion}}.0) Gecko/20100101 Firefox/{{.FirefoxVersion}}.0"
	if strings.Contains(platform, "Machintosh") {

		platform = strings.ReplaceAll(platform, "_", ".")
	}
	firefox, err := template.New("firefox").Parse(firefoxtemplate)
	if err != nil {
		return "", err
	}
	firefoxVersion := f.FirefoxVersions[random(len(f.FirefoxVersions))]

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
