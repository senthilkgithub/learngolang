package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filecopy("D:\\ng-book-AngularJS.pdf", "F:\\ng-book-AngularJS.pdf")
	fmt.Println("file has been copied successfully")
}

func filecopy(srcName, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	dest, err := os.Create(destName)
	if err != nil {
		return
	}
	defer dest.Close()
	wrbytes, err := io.Copy(dest, src)
	if err != nil {
		return
	}

	defer src.Close()
	fmt.Println(wrbytes)

	return
}
