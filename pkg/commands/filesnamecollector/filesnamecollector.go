package filesnamecollector

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/trypophob1a/fileschecker/pkg/config"
	"github.com/trypophob1a/fileschecker/pkg/core"
)

var (
	collect      = config.GetApp().Command("collect", "Collecting all file in the directory.")
	dir          = collect.Arg("dir", "scan directory.").Required().ExistingDir()
	saveRegistry = collect.Arg("registry", `Specify the path and file where to save the list of books
example: /home/usr/myBook.txt`).Required().String()
)

type CommandCollect struct {
	dir          string
	saveRegistry string
}

func NewCollector() *CommandCollect {
	return &CommandCollect{*dir, *saveRegistry}
}

func (s CommandCollect) getFilesName(scanPath string, extension []string) []string {
	files := make([]string, 0, 15)

	_ = filepath.Walk(scanPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && s.checkExtension(path, extension) {
			files = append(files, path)
		}

		return nil
	})

	return files
}

func (s CommandCollect) checkExtension(file string, extension []string) bool {
	for _, ex := range extension {
		if filepath.Ext(file) == strings.ToLower(ex) {
			return true
		}
	}

	return false
}

func (s CommandCollect) CollectInFile(scanPath, collectFile string, extension []string) {
	names := s.getFilesName(scanPath, extension)
	file := core.CreateFile(collectFile)

	defer file.Close()

	for _, name := range names {
		_, err := file.WriteString(name + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (s CommandCollect) isValidate() bool {
	match, _ := regexp.MatchString("^(.*[/\\\\])?[a-zA-zА-Яа-яЁё]+[.]txt", s.saveRegistry)
	return match
}

func (s CommandCollect) Execute() {
	if !s.isValidate() {
		println("error: argument 'registry' not a directory/file.txt, try --help")
		return
	}

	s.CollectInFile(s.dir, s.saveRegistry, config.Extensions())
}
