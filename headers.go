package fakeheaders

import (
	"embed"
	"encoding/json"
)

var (
	//go:embed static/*
	fs embed.FS
)

type FakeHeaders struct {
	UserAgents []string
}

func NewFakeHeaders() *FakeHeaders {
	//check if useragents.json exists
	//if not, download it
	//read useragents.json

	file, err := fs.ReadFile("static/useragents.json")

	if err != nil {
		downloadUserAgents()
		panic(err)
	}
	file, err = fs.ReadFile("static/useragents.json")
	//populate UserAgents slice
	var result []string
	json.Unmarshal([]byte(file), &result)
	return &FakeHeaders{
		UserAgents: result,
	}
}
