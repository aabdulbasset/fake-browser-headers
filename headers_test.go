package fakeheaders

import "testing"

func TestNewFakeHeaders(t *testing.T) {
	NewFakeHeaders(&FakeHeadersOptions{})
}

func TestFakeHeaders_RandomUserAgent(t *testing.T) {
	f := NewFakeHeaders(&FakeHeadersOptions{})
	_, err := f.RandomUserAgent()
	if err != nil {
		t.Error(err)
	}
}
func TestFakeHeaders_RandomAccept(t *testing.T) {
	f := NewFakeHeaders(&FakeHeadersOptions{})
	_, err := f.RandomAccept()
	if err != nil {
		t.Error(err)
	}
}

func TestFakeHeaders_RandomHeaders(t *testing.T) {
	f := NewFakeHeaders(&FakeHeadersOptions{BrowserToGenerate: Chrome})
	_, err := f.RandomHeaders()
	// println(fakeHeaders.UserAgent)
	if err != nil {
		t.Error(err)
	}
}
func TestFakeHeaders_FirefoxHeaders(t *testing.T) {
	f := NewFakeHeaders(&FakeHeadersOptions{BrowserToGenerate: Firefox})
	fakeHeaders, err := f.RandomHeaders()
	println(fakeHeaders.UserAgent)
	if err != nil {
		t.Error(err)
	}
}

func TestFakeHeaders_RandomAcceptLanguage(t *testing.T) {
	f := NewFakeHeaders(&FakeHeadersOptions{})
	_, err := f.RandomAcceptLanguage()

	if err != nil {
		t.Error(err)
	}
}
