package fakeheaders

import "testing"

func TestNewFakeHeaders(t *testing.T) {
	NewFakeHeaders()
}

func TestFakeHeaders_RandomUserAgent(t *testing.T) {
	f := NewFakeHeaders()
	_, err := f.RandomUserAgent()
	if err != nil {
		t.Error(err)
	}
}
