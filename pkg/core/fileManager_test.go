package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFile(t *testing.T) {
	err := CreateFile("sdsw:/#$@!wRs/file#$@!wRs.txt").Close()
	require.Error(t, err)

	err = CreateFile("./testdata/dir_test/file.txt").Close()

	if err != nil {
		return
	}

	isDirErr := CreateFile("./testdata/dir_test/").Close()
	require.Error(t, isDirErr)

	require.DirExists(t, "./testdata/dir_test")
	require.FileExists(t, "./testdata/dir_test/file.txt")
	_ = os.RemoveAll("./testdata/dir_test")
}

func TestCopyFile(t *testing.T) {
	_ = CopyFile("./testdata/TestExecutor.go", "./testdata/copy/TestExecutor.go")

	require.DirExists(t, "./testdata/copy")
	require.FileExists(t, "./testdata/copy/TestExecutor.go")

	err := CopyFile("./testdata/", "./testdata/copy/TestExecutor.go")
	require.Error(t, err)

	err = CopyFile("./testdata/TestExecutor.go", "sdsw:/#$@!wRs//testdata/copy/TestExecutor.go")
	require.Error(t, err)

	err = CopyFile("./testdata/TestExecutor.go", "./testdata/copy/")
	require.ErrorContains(t, err, "open ./testdata/copy/: is a directory")

	_ = os.RemoveAll("./testdata/copy")
}

func TestMoveFile(t *testing.T) {
	filePath := "./testdata/move_file_test!.txt"
	file := CreateFile(filePath)
	file.Close()
	_ = MoveFile(filePath, "./testdata/move/move_file_test!.txt")
	require.FileExists(t, "./testdata/move/move_file_test!.txt")
	require.DirExists(t, "./testdata/move")

	err := MoveFile("./testdata/", "./testdata/move/move_file_test!.txt")
	require.Error(t, err)

	file = CreateFile(filePath)
	err = MoveFile(filePath, "./testdata/move/move_file_test!.txt")
	require.Error(t, err)
	file.Close()
	_ = os.Remove(filePath)
	_ = os.RemoveAll("./testdata/move")
}

func TestGetFileName(t *testing.T) {
	require.Equal(t, "testexecutor", GetFileName("./testdata/TestExecutor.go"))
}

func TestUnSerializeTxt(t *testing.T) {
	filePath := "./testdata/test_file.txt"
	tests := []struct {
		path     string
		index    int
		expected string
	}{
		{path: filePath, index: 2, expected: "ccc"},
		{path: filePath, index: 0, expected: "aaa"},
		{path: filePath, index: 4, expected: "eee"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.path, func(t *testing.T) {
			result := UnSerializeTxt(tc.path, func(col []string, value string) []string {
				return append(col, value)
			}, make([]string, 0, 10))
			require.Equal(t, tc.expected, result[tc.index])
			require.Equal(t, 6, len(result))
		})
	}

	file := UnSerializeTxt("./testdata/test_file1.txt", func(col []string, value string) []string {
		return append(col, value)
	}, make([]string, 0, 10))

	if file == nil {
		require.Error(t, fmt.Errorf("can't open file"))
	}
}
