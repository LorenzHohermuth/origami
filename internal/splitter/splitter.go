package splitter

import (
	"fmt"
	"os"
	"strings"
)

func SplitFile(path string) RegistryEntry {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		paths := getFilePathsFromDir(path)
		arr := make([]RegistryEntry, len(paths))

		for _, filePath := range paths {
			if filePath != "" {
				arr = append(arr, SplitFile(filePath))
			}
		}
		return RegistryEntry{
			Name: info.Name(), 
			Tokens: nil,
			ChildEntrys: arr,
		}
	} else {
		return RegistryEntry{
			Name: info.Name(),
			Tokens: tokenize(getFileContent(path)),
			ChildEntrys: nil,
		}
	}
}

func getFileContent(path string) string{
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Content of File could not be found %s", path))
	}
	return string(data)
}

func getFilePathsFromDir(path string) []string{
	files, err := os.ReadDir(path)
	if err != nil {
		panic(fmt.Errorf("Directory could not be found %s", path))
	}
	var paths = make([]string, len(files))

	for _, file := range files {
		pathToFile := fmt.Sprintf("%s/%s", path, file.Name())
		paths = append(paths, pathToFile)
	}

	return paths
}

func tokenize(s string) []string {
	return strings.Fields(s)
}

func addToMatrix(addToMatrix [][]string, getsAddedToMatrix [][]string) [][]string {
	for  _, slice := range getsAddedToMatrix {
		addToMatrix = append(addToMatrix, slice)
	}
	return addToMatrix
}
