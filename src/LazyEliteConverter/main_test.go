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

func TestDefaultConfig(t *testing.T) {
	const (
		srcName = "srcTest"
		outputName = "outputTest")
	config := NewConfig(srcName, outputName)
	if config.srcFolder != srcName {
		t.Fatalf("config.srcFolder is not %v", srcName)
	}
	if config.outputFolder != outputName {
		t.Fatalf("config.outputFolder is not %v", outputName)
	}
	if config.Preserve != true {
		t.Fatal("config defaults does not preserve the source!!")
	}
	if config.Format != "png" {
		t.Fatal("config defaults does not use png")
	}
	fmt.Println(config)
}
