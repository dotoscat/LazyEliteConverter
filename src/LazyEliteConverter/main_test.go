package LazyEliteConverter

import (
	"fmt"
	"path/filepath"
	"testing"
)

var testbmp string = filepath.Join("testdata", "test.bmp")

func TestReadDir(t *testing.T) {
	expected := BitmapPaths{
		filepath.Join("testdata", "test.bmp")}
	files, err := GetBitmapList("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(files, err, expected)
}

func TestConvertBMPToPNG(t *testing.T) {
	outputFile := filepath.Join("testdata", "testtest.png")
	err := PathToPNGImage(testbmp, outputFile)
	if err != nil {
		t.Fatal(err)
	}
}
