package filesnamecollector

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var extension = []string{".pdf", ".doc", ".docx", ".txt", ".fb2", ".epub", ".djvu"}

func TestCommandCollect_getFilesName(t *testing.T) {
	formatPath := fmt.Sprintf("testdata%s", string(filepath.Separator))
	expect := []string{
		formatPath + "dddd.txt",
		formatPath + "err.txt",
		formatPath + "fffs.txt",
		formatPath + "helo wrld!.txt",
	}

	files := NewCollector().getFilesName("./testdata", extension)

	require.Equal(t, expect, files)
}

func TestCommandCollect_checkExtension(t *testing.T) {
	tests := []struct {
		file     string
		expected bool
	}{
		{file: "test_file.txt", expected: true},
		{file: "test_file.html", expected: false},
		{file: "test_file.pdf", expected: true},
		{file: "test_file.exe", expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.file, func(t *testing.T) {
			result := NewCollector().checkExtension(tc.file, extension)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCommandCollect_CollectInFile(t *testing.T) {
	NewCollector().CollectInFile("./testdata", "./testdata/uniq.txt", extension)

	// Checking creating file
	require.FileExists(t, "./testdata/uniq.txt")

	file, _ := os.Open("./testdata/uniq.txt")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		err = os.Remove("./testdata/uniq.txt")
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Checking file data
	formatPath := fmt.Sprintf("testdata%s", string(filepath.Separator))
	expect := []string{
		formatPath + "dddd.txt",
		formatPath + "err.txt",
		formatPath + "fffs.txt",
		formatPath + "helo wrld!.txt",
	}

	require.Equal(t, expect, lines)
}

func TestCommandCollect_isValidate(t *testing.T) {
	require.False(t, CommandCollect{"./testdata", "./testdata/uniq.tx"}.isValidate())
	require.True(t, CommandCollect{"./testdata", "./testdata/uniq.txt"}.isValidate())
}

func TestCommandCollect_Execute(t *testing.T) {
	collector := &CommandCollect{"./testdata", "./testdata/uniq.txt"}
	collector.Execute()
	require.FileExists(t, "./testdata/uniq.txt")
	_ = os.Remove("./testdata/uniq.txt")

	collector = &CommandCollect{"./testdata", "./testdata/uniq.xt"}
	collector.Execute()
	require.NoFileExists(t, "./testdata/uniq.txt")
}
