package fakeheaders

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func randomfloat() float64 {
	return math.Min(rand.Float64()+0.2, 1)
}

func (f *FakeHeaders) RandomAcceptLanguage() (string, error) {
	chosenBefore := make(map[string]bool)
	if len(f.AcceptLanguages) <= 0 {
		return "", errors.New("No AcceptLanguages found")
	}
	uaString := "en-US"
	for i := 0; i < rand.Intn(5)+1; i++ {
		chosen := f.AcceptLanguages[random(len(f.AcceptLanguages))]
		if chosenBefore[chosen] {
			continue
		}
		chosenBefore[chosen] = true
		if strings.Contains(chosen, "-") {
			uaString += "," + chosen
		} else {
			uaString += "," + chosen + fmt.Sprintf(";q=%.1f", randomfloat())
		}
	}
	return uaString, nil
}
