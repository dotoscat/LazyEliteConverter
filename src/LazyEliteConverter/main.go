package LazyEliteConverter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"errors"

	"golang.org/x/image/bmp"
	"image/png"
)

// Config describes the config for the conversion
type Config struct {
	Preserve     bool
	Format       string
	srcFolder    string
	outputFolder string
}

func (c Config) SrcFolder() string {
	return c.srcFolder
}

func (c Config) OutputFolder() string {
	return c.outputFolder
}

type (
	BitmapPaths []string
	Paths       []string
)

// NewConfig returns a Config with some of the defaults
func NewConfig(srcFolder, outputFolder string) Config {
	return Config{
		true,
		"png",
		srcFolder,
		outputFolder}
}

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

func GetOutputList(list BitmapPaths, config Config) Paths {
	paths := make(Paths, 0)
	for _, path := range list {
		base := filepath.Base(path)
		chunks := strings.Split(base, ".")
		outputName := chunks[0] + "." + config.Format
		outputPath := filepath.Join(config.OutputFolder(), outputName)
		paths = append(paths, outputPath)
	}
	return paths
}

// This transforms a bmp image from a souce to a png into an output
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

func ConvertList(config Config) error {
	list, err := GetBitmapList(config.SrcFolder())
	if err != nil {
		return err
	}
	fmt.Println(list)
	outputList := GetOutputList(list, config)
	if len(list) != len(outputList) {
		return errors.New("For some reason the len of the bitmap list is not the same as the output list")
	}
	for i := 0; i < len(list); i++ {
		err := PathToPNGImage(list[i], outputList[i])
		if err != nil {
			return err
		}
		if !config.Preserve {
			fmt.Printf("Not preserve %v\n", list[i])
		}
	}
	return nil
}
