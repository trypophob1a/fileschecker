package strategy

import (
	fuzzy "github.com/reallyliri/go-fuzzywuzzy"

	"github.com/trypophob1a/fileschecker/pkg/core"
)

type DefaultFinder struct {
	firstTxt, secondTxt string
}

func NewDefaultFinder() *DefaultFinder {
	return &DefaultFinder{}
}

// Find TODO optimize the algorithm.
func (f *DefaultFinder) Find(percent uint8, callback func(filename string)) {
	firstList := f.unSerializeTxt(f.firstTxt)
	secondList := f.unSerializeTxt(f.secondTxt)

	for _, s := range secondList {
		if f.contains(firstList, s, percent) {
			continue
		}

		callback(s)
	}
}

func (f *DefaultFinder) SetResources(first, second string) {
	f.firstTxt = first
	f.secondTxt = second
}

func (f DefaultFinder) contains(slice []string, item string, percent uint8) bool {
	for _, el := range slice {
		if uint8(fuzzy.Ratio(core.GetFileName(el), core.GetFileName(item))) >= percent {
			return true
		}
	}

	return false
}

func (f DefaultFinder) unSerializeTxt(path string) []string {
	return core.UnSerializeTxt(path, func(col []string, value string) []string {
		return append(col, value)
	}, make([]string, 0, 10))
}
