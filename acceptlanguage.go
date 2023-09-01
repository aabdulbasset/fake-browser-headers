package fakeheaders

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func (f *FakeHeaders) RandomAcceptLanguage() (string, error) {
	if len(f.AcceptLanguages) <= 0 {
		return "", errors.New("No AcceptLanguages found")
	}
	uaString := f.AcceptLanguages[random(len(f.AcceptLanguages))]
	chosenBefore := make(map[string]bool)
	chosenBefore[uaString] = true
	quality := 0.9
	for i := 0; i < rand.Intn(3)+1; i++ {
		chosen := f.AcceptLanguages[random(len(f.AcceptLanguages))]
		if chosenBefore[chosen] {
			continue
		}
		chosenBefore[chosen] = true
		uaString += "," + chosen + fmt.Sprintf(";q=%.1f", quality)
		quality = math.Max(0.1, quality-0.1)
	}
	return uaString, nil
}
