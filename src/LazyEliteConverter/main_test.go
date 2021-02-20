package LazyEliteConverter

import (
	"fmt"
	"path/filepath"
	"testing"
)

const (
	srcName    = "srcTest"
	outputName = "outputTest"
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

func TestGetOutputList(t *testing.T) {
	config := NewConfig("./testdata", "./testdata")
	srcFiles, err := GetBitmapList(config.SrcFolder())
	if err != nil {
		t.Fatal(err)
	}
	outputFiles := GetOutputList(srcFiles, config)
	fmt.Println(outputFiles)
}

func TestConvertBMPToPNG(t *testing.T) {
	outputFile := filepath.Join("testdata", "testtest.png")
	err := PathToPNGImage(testbmp, outputFile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDefaultConfig(t *testing.T) {
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

func TestConvertList(t *testing.T){
	config := NewConfig("./testdata", "./testdata")
	err := ConvertList(config)
	if err != nil {
		t.Fatal(err)
	}
}
