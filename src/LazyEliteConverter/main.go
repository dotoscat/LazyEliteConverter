package LazyEliteConverter

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type BitmapPaths []string

func GetBitmapList(srcFolder string) (BitmapPaths, error) {
	filesInfo, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		return nil, err
	}
	paths := make(BitmapPaths, 0)
	for _, fileInfo := range filesInfo {
		isBitmap := strings.HasSuffix(strings.ToLower(fileInfo.Name()), ".bmp")
		if fileInfo.IsDir() || !isBitmap {
			continue
		}
		path := filepath.Join(srcFolder, fileInfo.Name())
		paths = append(paths, path)
	}
	return paths, nil
}

func Algo() {
	fmt.Println("Algo")
}
