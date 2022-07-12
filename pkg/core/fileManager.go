package core

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/trypophob1a/fileschecker/pkg/interfaces"
)

func CreateFile(filePath string) *os.File {
	dir, fileName := filepath.Split(filePath)
	if err := os.MkdirAll(dir, 0o600); err != nil {
		fmt.Println(err)
		return nil
	}

	newFile, err := os.Create(dir + string(filepath.Separator) + fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return newFile
}

func CopyFile(from, to string) error {
	dir := filepath.Dir(to)

	bytesRead, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, 0o777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(to, bytesRead, 0o600)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetFileName(filename string) string {
	_, name := filepath.Split(filename)
	return strings.ToLower(strings.TrimSuffix(name, path.Ext(filename)))
}

func UnSerializeTxt[T interfaces.Collection](path string, adder func(col T, value string) T, coll T) T {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		coll = adder(coll, scanner.Text())
	}

	return coll
}
