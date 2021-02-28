// LazyEliteConverter - Convert bmp files into lossless compressed images
// Copyright (C) 2021  Oscar Triano García

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package main

import "github.com/dotoscat/LazyEliteConverter/src/LazyEliteConverter"
import "flag"
import "fmt"
import "os"

const sourceUsage = "Source directory"
const outputUsage = "Output directory"
const noPreserveFilesUsage = "No preserve files"

const VERSION = "0.1"

func main() {
	var source string
	var output string
	var noPreserveFiles bool = false
	var showVersion = false
	flag.StringVar(&source, "source", "", sourceUsage)
	flag.StringVar(&source, "s", "", sourceUsage)
	flag.StringVar(&output, "output", "", outputUsage)
	flag.StringVar(&output, "o", "", outputUsage)
	flag.BoolVar(&noPreserveFiles, "no-preserve-original-files", false, noPreserveFilesUsage)
	flag.BoolVar(&showVersion, "version", false, "")
	flag.BoolVar(&showVersion, "v", false, "")
	flag.Parse()
	if showVersion{
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if len(source) == 0 {
		fmt.Println("source is empty!")
		os.Exit(1)
	}
	if len(output) == 0 {
		fmt.Println("output is empty!")
		os.Exit(2)
	}
	fmt.Printf("Convert bmp files from \"%v\" to \"%v\"\n", source, output)
	config := LazyEliteConverter.NewConfig(source, output)
	config.Preserve = !noPreserveFiles
	fmt.Println("Preserve files? ", config.Preserve)
	err := LazyEliteConverter.ConvertList(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
