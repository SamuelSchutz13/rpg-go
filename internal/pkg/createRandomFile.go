package pkg

import (
	"os"
	"strconv"
)

func CreateRandomFile(LocalPath string, filename string) (string, error) {
	files, err := os.ReadDir(LocalPath)

	if err != nil {
		panic(err)
	}

	for _, v := range files {
		if v.Name() == filename {
			newFilename := filename + strconv.Itoa(len(files)+1)
			return newFilename, nil
		}
	}

	return filename, nil
}
