package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(ioutil.ReadDir("."))
	fmt.Println(os.Getwd())
	fpath, _ := filepath.Abs("./")
	destPath := fmt.Sprintf(fpath + "\\report")
	fmt.Println(destPath)

}
