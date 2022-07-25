package selfcheck

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/trypophob1a/fileschecker/pkg/config"
	"github.com/trypophob1a/fileschecker/pkg/core"
	"github.com/trypophob1a/fileschecker/pkg/interfaces"
	"github.com/trypophob1a/fileschecker/pkg/strategy/selfcheckfinder"
)

var (
	check = config.GetApp().Command("selfcheck", "self check for duplicate files.")
	file  = check.Arg("file", "file. example: /home/usr/File.txt").Required().
		ExistingFile()
	percent = check.Flag("percent", "how many percent similarity: default = 85").Short('p').
		Default("85").Uint8()
)

type CommandSelfCheck struct {
	file    string
	percent uint8
	finder  interfaces.SelfCheckFinder
}

func NewSelfCheck() *CommandSelfCheck {
	return &CommandSelfCheck{file: *file, percent: *percent}
}

func (c *CommandSelfCheck) SetFinder(finder interfaces.SelfCheckFinder) *CommandSelfCheck {
	c.finder = finder
	c.finder.SetResource(c.file)
	return c
}

func (c CommandSelfCheck) Check(percent uint8, callback func(filename string)) {
	if c.finder == nil {
		c.SetFinder(selfcheckfinder.NewDefaultFinder()).finder.Find(percent, callback)
		return
	}

	c.finder.Find(percent, callback)
}

func (c CommandSelfCheck) Execute() {
	separator := string(filepath.Separator)
	pathFile := filepath.Dir(c.file)
	file := core.CreateFile(pathFile + separator + "duplicate_files.txt")
	defer file.Close()

	c.Check(c.percent, func(duplicateFilename string) {
		absPath, err := filepath.Abs(duplicateFilename)
		if err != nil {
			fmt.Println(err)
			return
		}

		copyPath := pathFile + separator + "duplicate_files" + strings.TrimPrefix(absPath, pathFile)
		err = core.MoveFile(duplicateFilename, copyPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.WriteString(duplicateFilename + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
