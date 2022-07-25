package selfcheckfinder

import (
	fuzzy "github.com/reallyliri/go-fuzzywuzzy"

	"github.com/trypophob1a/fileschecker/pkg/core"
)

type DefaultFinder struct {
	resourcePath string
}

func NewDefaultFinder() *DefaultFinder {
	return &DefaultFinder{}
}

func (f DefaultFinder) Find(percent uint8, callback func(duplicateFilename string)) {
	fileNames := f.unSerializeTxt()

	for i := 0; i < len(fileNames); i++ {
		for j := 1 + i; j < len(fileNames); j++ {
			if uint8(fuzzy.Ratio(core.GetFileName(fileNames[i]), core.GetFileName(fileNames[j]))) >= percent {
				callback(fileNames[j])
			}
		}
	}
}

func (f *DefaultFinder) SetResource(fileTxt string) {
	f.resourcePath = fileTxt
}

func (f DefaultFinder) unSerializeTxt() []string {
	return core.UnSerializeTxt(f.resourcePath, func(collection []string, value string) []string {
		return append(collection, value)
	}, make([]string, 0, 50))
}

func hello() {
	println("hello")
}
