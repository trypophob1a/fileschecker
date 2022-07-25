package check

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trypophob1a/fileschecker/pkg/strategy/checkfinder"
)

func TestNewCheck(t *testing.T) {
	checker := NewCheck()
	require.IsType(t, &CommandCheck{}, checker)
}

func TestCommandCheck_Check(t *testing.T) {
	checker := &CommandCheck{first: "./testdata/first.txt", second: "./testdata/second.txt", percent: 90}
	expect := []string{
		"./testdata/files/dddd.txt",
		"./testdata/files/err.txt",
		"./testdata/files/fffs.txt",
	}

	actual := make([]string, 0)
	checker.Check(checker.percent, func(filename string) {
		actual = append(actual, filename)
	})
	sort.Strings(actual)
	require.Equal(t, expect, actual)

	checker.SetFinder(checkfinder.NewHashmapFinder())
	actual = make([]string, 0)
	checker.Check(checker.percent, func(filename string) {
		actual = append(actual, filename)
	})
	sort.Strings(actual)
	require.Equal(t, expect, actual)
}

func TestCommandCheck_Execute(t *testing.T) {
	first, _ := filepath.Abs("./testdata/first.txt")
	second, _ := filepath.Abs("./testdata/second.txt")
	sep := string(filepath.Separator)
	checker := CommandCheck{first: first, second: second, percent: 90}
	formatPath := fmt.Sprintf("testdata%suniq_files%sfiles%s", sep, sep, sep)
	expect := []string{
		formatPath + "dddd.txt",
		formatPath + "err.txt",
		formatPath + "fffs.txt",
	}

	checker.Execute()
	require.DirExists(t, "./testdata/uniq_files")
	require.FileExists(t, "./testdata/uniq_files.txt")

	files := make([]string, 0, 3)

	e := filepath.Walk("./testdata/uniq_files/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if e != nil {
		log.Fatal(e)
	}

	require.Equal(t, expect, files)

	file, _ := os.Open("./testdata/uniq_files.txt")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		e = os.Remove("./testdata/uniq_files.txt")
		if e != nil {
			log.Fatal(e)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	require.Equal(t, []string{
		"./testdata/files/dddd.txt",
		"./testdata/files/fffs.txt",
		"./testdata/files/err.txt",
	}, lines)
	_ = os.RemoveAll("./testdata/uniq_files/")

	// testing errors
	first, _ = filepath.Abs("./testdata/first.txt")
	second, _ = filepath.Abs("./testdata/second_with_error_path_for_copy.txt")
	checker = CommandCheck{first: first, second: second, percent: 90}
	checker.Execute()
	require.NoFileExists(t, "./testdata/uniq_files/file/1.txt")
	_ = os.RemoveAll("./testdata/uniq_files/")

	first, _ = filepath.Abs("./testdata/first.txt")
	checker = CommandCheck{first: first, second: "./testdata/second_with_error_path_for_Abs.txt", percent: 90}
	checker.Execute()
	require.NoDirExists(t, "./testdata/uniq_files")
	_ = os.RemoveAll("./testdata/uniq_files/")
}
