package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bangnh1/golang-training/06/cfg"
	"github.com/spf13/viper"
)

// List all files in the upload directory
func ListAllFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := GetFileName(path)
			if err != nil {
				return nil
			}
			files = append(files, file)
		}
		return nil
	})
	return files, err
}

// Read log file to get processed files
func ReadLogFile() []string {
	bytesRead, err := ioutil.ReadFile(viper.GetString(cfg.ConfigKeyLogInfoFileName))
	if err != nil {
		return []string{}
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	return lines
}

// Get filename and check file type
func GetFileName(filePath string) (string, error) {
	pathLevel := strings.Split(filePath, "/")
	filename := pathLevel[len(pathLevel)-1]
	if fileExtension := filepath.Ext(filename); fileExtension != ".mp4" {
		return "", fmt.Errorf("file extension ins't equal to .mp4")
	}
	filenameWithoutExt := strings.Split(filename, ".mp4")[0]

	return filenameWithoutExt, nil
}

// Get unhashed filename
func GetHashingFiles(uploadedFiles []string, loggedFiles []string) []string {
	var diff []string

	for _, file := range uploadedFiles {
		con := contains(loggedFiles, file)
		if !con {
			diff = append(diff, file)
		}
	}
	return diff
}

// check if string array contains string or not
func contains(s []string, searchterm string) bool {
	for _, file := range s {
		if file == searchterm {
			return true
		}
	}
	return false
}
