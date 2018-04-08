package main

import (
	"fmt"
	"os"
)

func main() {
	fname := os.Args[:1]
	if len(fname) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: png_parse [png_file]¥n")
	}
	pngData, err := getPngImage(fname[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s¥n", err)
	}
	printPngInfo(pngData)
}

func getPngImage(fileName string) (PNGData, error) {
	fmt.Fprintln(os.Stderr, "Enter getPngImage¥r¥n")
	data := PNGData{}

	return data, nil
}

func printPngInfo(data PNGData) {
	fmt.Fprintln(os.Stderr, "Enter printPngInfo¥r¥n")
}

type PNGData struct {
	header PNGHeader
	data   PNGImage
}

type PNGHeader struct {
	length      int32
	width       int32
	height      int32
	depth       byte
	colorType   byte
	compression byte
	filter      byte
	interlace   byte
}

type PNGImage struct {
	length int32
	data   []int32
}
