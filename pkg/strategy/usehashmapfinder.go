package strategy

import (
	fuzzy "github.com/reallyliri/go-fuzzywuzzy"

	"github.com/trypophob1a/fileschecker/pkg/core"
)

type HashmapFinder struct {
	firstTxt, secondTxt string
}

func NewHashmapFinder() *HashmapFinder {
	return &HashmapFinder{}
}

func (f HashmapFinder) contains(hashmap map[string]struct{}, item string, percent uint8) bool {
	for key := range hashmap {
		if uint8(fuzzy.Ratio(core.GetFileName(key), core.GetFileName(item))) >= percent {
			return true
		}
	}

	return false
}

func (f HashmapFinder) unSerializeTxt(path string) map[string]struct{} {
	return core.UnSerializeTxt(path, func(col map[string]struct{}, value string) map[string]struct{} {
		col[value] = struct{}{}
		return col
	}, make(map[string]struct{}, 10))
}

func (f *HashmapFinder) SetResources(first, second string) {
	f.firstTxt = first
	f.secondTxt = second
}

func (f HashmapFinder) Find(percent uint8, callback func(filename string)) {
	first := f.unSerializeTxt(f.firstTxt)
	second := f.unSerializeTxt(f.secondTxt)

	for key := range second {
		if _, ok := first[core.GetFileName(key)]; ok {
			continue
		}

		if f.contains(first, key, percent) {
			continue
		}

		first[key] = struct{}{}
		callback(key)
	}
}
