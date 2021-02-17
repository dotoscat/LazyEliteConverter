package LazyEliteConverter

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
	
	"golang.org/x/image/bmp"
	"image/png"
)

type BitmapPaths []string

// GetBitmapList gets a list of string with the the BMP paths
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

func PathToPNGImage(src, output string) error {
	srcFile, err := os.Open(src)
	defer srcFile.Close()
	if err != nil {
		return err
	}
	bmp, err := bmp.Decode(srcFile)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		return err
	}
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	if err := encoder.Encode(outputFile, bmp); err != nil {
		return err
	}
	return nil
}

func Algo() {
	fmt.Println("Algo")
}
