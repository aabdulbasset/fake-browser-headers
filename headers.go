package fakeheaders

import (
	"crypto/rand"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/mileusna/useragent"
)

var (
	//go:embed static/*
	fs embed.FS
)

type FakeHeaders struct {
	UserAgents      []string
	Accepts         []string
	AcceptLanguages []string
	AcceptEncodings []string
}
type FakeHeader struct {
	UserAgent               string
	Accept                  string
	AcceptLanguage          string
	AcceptEncoding          string
	Connection              string
	UpgradeInsecureRequests string
	SecFetchUser            string
	SecFetchSite            string
	SecFetchMode            string
	SecFetchDest            string
	SecFetchPlatform        string
	SecMobile               string
	SecUA                   string
	Te                      string
}

func NewFakeHeaders() *FakeHeaders {
	//check if useragents.json exists
	//if not, download it
	//read useragents.json

	userAgents, err := fs.ReadFile("static/useragents.json")

	if err != nil {
		downloadUserAgents()
		panic(err)
	}
	userAgents, _ = fs.ReadFile("static/useragents.json")
	//populate UserAgents slice
	var userAgentsResult []string
	json.Unmarshal([]byte(userAgents), &userAgentsResult)

	acceptFile, err := fs.ReadFile("static/accept.json")
	if err != nil {
		panic(err)
	}
	var acceptResult []string
	json.Unmarshal([]byte(acceptFile), &acceptResult)

	acceptLanguageFile, err := fs.ReadFile("static/acceptlanguage.json")
	if err != nil {
		panic(err)
	}
	var acceptLanguageResult []string
	json.Unmarshal([]byte(acceptLanguageFile), &acceptLanguageResult)

	acceptEncodingFile, err := fs.ReadFile("static/acceptencode.json")
	if err != nil {
		panic(err)
	}
	var acceptEncodingResult []string
	json.Unmarshal([]byte(acceptEncodingFile), &acceptEncodingResult)

	return &FakeHeaders{
		UserAgents:      userAgentsResult,
		Accepts:         acceptResult,
		AcceptLanguages: acceptLanguageResult,
		AcceptEncodings: acceptEncodingResult,
	}
}

func (f *FakeHeaders) RandomAccept() (string, error) {
	if len(f.Accepts) <= 0 {
		return "", errors.New("No Accepts found")
	}
	return f.Accepts[random(len(f.Accepts))], nil
}

func (f *FakeHeaders) RandomAcceptEncoding() (string, error) {
	if len(f.AcceptEncodings) <= 0 {
		return "", errors.New("No AcceptEncodings found")
	}
	return f.AcceptEncodings[random(len(f.AcceptEncodings))], nil
}

func (f *FakeHeaders) RandomHeaders() (*FakeHeader, error) {
	randUserAgent, err := f.RandomUserAgent()
	randAccept, err := f.RandomAccept()
	randAcceptLanguage, err := f.RandomAcceptLanguage()
	randAcceptEncoding, err := f.RandomAcceptEncoding()
	if err != nil {
		return &FakeHeader{}, err
	}
	ua := useragent.Parse(randUserAgent)

	derivedUA := ""
	SecPlatform := ""
	SecMobile := ""
	te := ""

	if strings.Contains(randUserAgent, "Firefox") {
		te = "trailers"
	}
	if strings.Contains(randUserAgent, "Chrome") {
		derivedUA = fmt.Sprintf("\"Chromium\";v=\"%d\", \"Not A Brand\";v=\"%d\"", ua.VersionNo.Major, random(99))
		if ua.Name == "Edge" {
			derivedUA += fmt.Sprintf(", \"Microsoft Edge\";v=\"%d\"", ua.VersionNo.Major)
		} else {
			derivedUA += fmt.Sprintf(", \"Google Chrome\";v=\"%d\"", ua.VersionNo.Major)
		}
		if ua.Mobile {
			SecMobile = "?1"
		} else {
			SecMobile = "?0"
		}
		SecPlatform = fmt.Sprintf("\"%s\"", ua.OS)
		//
	}

	return &FakeHeader{
		UserAgent:               randUserAgent,
		Accept:                  randAccept,
		AcceptLanguage:          randAcceptLanguage,
		Connection:              "keep-alive",
		UpgradeInsecureRequests: "1",
		SecFetchUser:            "?1",
		SecFetchSite:            "none",
		SecFetchMode:            "navigate",
		SecFetchDest:            "document",
		SecFetchPlatform:        SecPlatform, //TODO: Detect platform from userAgent
		SecMobile:               SecMobile,   //TODO: Detect if mobile from userAgent
		SecUA:                   derivedUA,
		AcceptEncoding:          randAcceptEncoding,
		Te:                      te,
	}, nil
}

func random(max int) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return n.Int64()

}
