package splitter

import (
	"fmt"
	"os"
	"strings"
)

func SplitFile(path string) [][]string {
	if path == "" {
		return make([][]string, 0)
	}
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		paths := getFilePathsFromDir(path)
		matrix := make([][]string, len(paths))
		for _, filePath := range paths {
			matrix = addToMatrix(matrix, SplitFile(filePath))
		}
		return matrix
	} else {
		matrix := make([][]string, 1)
		matrix[0] = tokenize(getFileContent(path))
		return matrix
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
