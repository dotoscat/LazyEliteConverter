package LazyEliteConverter

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	expected := BitmapPaths{
		filepath.Join("testdata", "test.bmp")}
	files, err := GetBitmapList("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(files, err, expected)
}
