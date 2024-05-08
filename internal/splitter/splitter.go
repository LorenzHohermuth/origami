package splitter

import (
	"fmt"
	"os"
	"strings"
)

func SplitFile(path string) []string {
	fileContent := getFileContent(path)
	return strings.Fields(fileContent)
}

func getFileContent(path string) string{
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Content of File could not be found %s", path))
	}
	return string(data)
}
