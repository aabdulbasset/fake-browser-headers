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
func TestFakeHeaders_RandomAccept(t *testing.T) {
	f := NewFakeHeaders()
	_, err := f.RandomAccept()
	if err != nil {
		t.Error(err)
	}
}

func TestFakeHeaders_RandomHeaders(t *testing.T) {
	f := NewFakeHeaders()
	fakeHeaders, err := f.RandomHeaders()
	println(fakeHeaders.Connection)
	if err != nil {
		t.Error(err)
	}
}

func TestFakeHeaders_RandomAcceptLanguage(t *testing.T) {
	f := NewFakeHeaders()
	_, err := f.RandomAcceptLanguage()

	if err != nil {
		t.Error(err)
	}
}
