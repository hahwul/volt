package file

import (
	"fmt"
	"os"
)

func GetFiles(path string) ([]string, error) {
	var fileList []string
	file, err := os.Open(path)
	if err != nil {
		return fileList, err
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		file, err := os.Open(filePath)
		if err != nil {
			return fileList, err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return fileList, err
		}
		fileList = append(fileList, filePath)
		if fileInfo.IsDir() {
			readDir(filePath)
		}
	}
	return fileList, nil
}
