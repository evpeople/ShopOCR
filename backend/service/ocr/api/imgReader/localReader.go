package imgReader

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type LocalReader struct {
	filePath []string
	index    int
}

func NewLocalReader(root string) LocalReader {
	imagePaths := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
			imagePaths = append(imagePaths, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	return LocalReader{filePath: imagePaths, index: 0}
}
func (lr *LocalReader) ReadImg() (base64Img string, err error) {
	path := lr.filePath[lr.index]
	if lr.index < len(lr.filePath) {
		lr.index++
	} else {
		return "", errors.New("no more image")
	}
	imageData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	return base64.StdEncoding.EncodeToString(imageData), nil
}
