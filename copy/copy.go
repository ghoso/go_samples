package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: copy.go [src] [dst]")
		os.Exit(1)
	}
	src := os.Args[1]
	dst := os.Args[2]
	if src == dst {
		fmt.Println("src and dst file is identical")
		os.Exit(1)
	}
	ret := copyFile(src, dst)
	if ret != nil {
		fmt.Printf("Error: %s\n", ret.Error())
	}
}

func copyFile(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}
