package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	src := "./src.tmp"
	dst := "./dst.tmp"
	fdata := "Hello World !"

	srcf, err := os.OpenFile(src, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer srcf.Close()
	fmt.Fprint(srcf, fdata)

	err = copyFile(src, dst)
	data, err := ioutil.ReadFile(dst)
	if string(data) != fdata {
		t.Error("did not match file contents")
	}

	os.Remove(src)
	os.Remove(dst)
}
