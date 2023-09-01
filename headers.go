package fakeheaders

import (
	"crypto/rand"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"strings"

	"github.com/mileusna/useragent"
)

var (
	//go:embed static/*
	fs embed.FS
)

type FakeHeaders struct {
	Accepts         []string
	AcceptLanguages []string
	AcceptEncodings []string
	Platforms       []string
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
	platforms := []string{
		"Windows NT 10.0; Win64; x64",
		"Windows NT 10.0; WOW64",
		"Windows NT 6.3; Win64; x64",
		"Windows NT 6.3; WOW64",
		"Machintosh; Intel Mac OS X 10_15_7",
		"Machintosh; Intel Mac OS X 10_15_6",
		"Machintosh; Intel Mac OS X 10_10",
		"Machintosh; Intel Mac OS X 10_9",
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux x86_64",
		"X11; Fedora; Linux x86_64",
		"X11; Arch; Linux x86_64",
		"X11; Gentoo; Linux x86_64",
	}

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
		Platforms:       platforms,
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
	punctuation := []string{".", "/", ")", "(", "_", " ", ":"}

	if err != nil {
		return &FakeHeader{}, err
	}
	ua := useragent.Parse(randUserAgent)

	SecPlatform := ""
	SecMobile := ""
	te := ""
	derivedUA := []string{}
	if strings.Contains(randUserAgent, "Firefox") {
		te = "trailers"
	}
	if strings.Contains(randUserAgent, "Chrome") {
		derivedUA = []string{
			fmt.Sprintf("\"Chromium\";v=\"%d\"", ua.VersionNo.Major),
		}
		if ua.Name == "Edge" {
			derivedUA = append(derivedUA, fmt.Sprintf("\"Microsoft Edge\";v=\"%d\"", ua.VersionNo.Major))
		} else {
			derivedUA = append(derivedUA, fmt.Sprintf("\"Google Chrome\";v=\"%d\"", ua.VersionNo.Major))
		}
		if ua.Mobile {
			SecMobile = "?1"
		} else {
			SecMobile = "?0"
		}
		SecPlatform = fmt.Sprintf("\"%s\"", ua.OS)
		//

		derivedUA = append(derivedUA, fmt.Sprintf("\"Not%sA%sBrand\";v=\"99\"", punctuation[random(len(punctuation))], punctuation[random(len(punctuation))]))
		mathrand.Shuffle(len(derivedUA), func(i, j int) { derivedUA[i], derivedUA[j] = derivedUA[j], derivedUA[i] })
	}
	var finalUA string
	if len(derivedUA) > 0 {
		finalUA = strings.Join(derivedUA, ", ")
	} else {
		finalUA = ""
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
		SecUA:                   finalUA,
		AcceptEncoding:          randAcceptEncoding,
		Te:                      te,
	}, nil
}

func random(max int) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return n.Int64()

}
