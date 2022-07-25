package selfcheck

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trypophob1a/fileschecker/pkg/core"
	"github.com/trypophob1a/fileschecker/pkg/strategy/selfcheckfinder"
)

func TestNewSelfCheck(t *testing.T) {
	checker := &CommandSelfCheck{}
	require.IsType(t, checker, NewSelfCheck())
}

func TestCommandSelfCheck_Check(t *testing.T) {
	checker := &CommandSelfCheck{file: "./testdata/second.txt", percent: 85}
	expect := []string{
		"./testdata/files/hello world!.txt",
		"./testdata/files/all_errirs.txt",
		"./testdata/files/all_erors.txt",
	}

	actual := make([]string, 0)
	checker.Check(checker.percent, func(filename string) {
		actual = append(actual, filename)
	})
	require.Equal(t, expect, actual)

	checker.SetFinder(selfcheckfinder.NewDefaultFinder())
	actual = make([]string, 0)
	checker.Check(checker.percent, func(filename string) {
		actual = append(actual, filename)
	})
	require.Equal(t, expect, actual)
}

func TestCommandSelfCheck_Execute(t *testing.T) {
	second, _ := filepath.Abs("./testdata/second.txt")
	files := []string{"all_erors.txt", "all_errirs.txt", "hello world!.txt"}
	checker := &CommandSelfCheck{file: second, percent: 85}
	checker.Execute()
	require.DirExists(t, "./testdata/duplicate_files")
	require.FileExists(t, "./testdata/duplicate_files.txt")
	for _, file := range files {
		require.FileExists(t, "./testdata/duplicate_files/files/"+file)
		err := core.MoveFile("./testdata/duplicate_files/files/"+file, "./testdata/files/"+file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	_ = os.RemoveAll("./testdata/duplicate_files/")

	duplicateFiles, _ := os.Open("./testdata/duplicate_files.txt")

	var lines []string
	scanner := bufio.NewScanner(duplicateFiles)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	require.Equal(t, []string{
		"./testdata/files/hello world!.txt",
		"./testdata/files/all_errirs.txt", "./testdata/files/all_erors.txt",
	}, lines)

	duplicateFiles.Close()

	// testing errors
	checker = &CommandSelfCheck{file: "./testdata/second_with_error_path_for_move.txt", percent: 85}
	checker.Execute()
	require.NoDirExists(t, "./testdata/duplicate_files")
	checker = &CommandSelfCheck{file: "./testdata/second_with_error_path_for_Abs.txt", percent: 85}
	checker.Execute()
	require.NoDirExists(t, "./testdata/duplicate_files")
	_ = os.Remove("./testdata/duplicate_files.txt")
}
