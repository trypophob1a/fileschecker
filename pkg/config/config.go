package config

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var Extensions = func() []string {
	return []string{".pdf", ".doc", ".docx", ".txt", ".fb2", ".epub", ".djvu"}
}

var singleInstance *kingpin.Application

func GetApp() *kingpin.Application {
	if singleInstance == nil {
		singleInstance = kingpin.New("Files checker", "A command-line checking duplicate files application.")

		return singleInstance
	}

	return singleInstance
}
