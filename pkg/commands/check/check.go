package check

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/trypophob1a/fileschecker/pkg/config"
	"github.com/trypophob1a/fileschecker/pkg/core"
	"github.com/trypophob1a/fileschecker/pkg/interfaces"
	"github.com/trypophob1a/fileschecker/pkg/strategy"
)

var (
	check = config.GetApp().Command("check", "Check for duplicate files.")
	first = check.Arg("first", "first file. example: /home/usr/firstFiles.txt").Required().
		ExistingFile()
	second = check.Arg("second", "second file. example: /home/usr/secondFiles.txt").Required().
		ExistingFile()
	percent = check.Flag("percent", "how many percent similarity: default = 90").Short('p').
		Default("90").Uint8()
)

type CommandCheck struct {
	first   string
	second  string
	percent uint8
	finder  interfaces.Finder
}

func NewCheck() *CommandCheck {
	return &CommandCheck{first: *first, second: *second, percent: *percent}
}

func (c *CommandCheck) SetFinder(finder interfaces.Finder) *CommandCheck {
	c.finder = finder
	c.finder.SetResources(c.first, c.second)
	return c
}

func (c CommandCheck) Check(percent uint8, callback func(filename string)) {
	if c.finder == nil {
		c.SetFinder(strategy.NewDefaultFinder()).finder.Find(percent, callback)
		return
	}

	c.finder.Find(percent, callback)
}

func (c CommandCheck) Execute() {
	separator := string(filepath.Separator)
	pathFile := filepath.Dir(c.second)
	file := core.CreateFile(pathFile + separator + "uniq_files.txt")

	defer file.Close()

	c.Check(c.percent, func(filename string) {
		absPath, err := filepath.Abs(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		copyPath := pathFile + separator + "uniq_files" + strings.TrimPrefix(absPath, pathFile)
		err = core.CopyFile(filename, copyPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.WriteString(filename + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
