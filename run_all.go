package main

import (
	"ccpAPITest/public/runTestCase"
	"fmt"
	"log"
	"os"
	"runtime"
	//"log"
	//"os"
)

func runLevel(args ...string) {
	fmt.Println(args)
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get basePath")
	}
	fmt.Println(len(args))
	if len(args) != 0 {
		for _, p := range args {
			destPath := ""
			if runtime.GOOS == "windows" {
				destPath = fmt.Sprintf(basePath+"\\testcase\\test_%v", p)
				fmt.Println("Test123", destPath)
			} else {
				destPath = fmt.Sprintf(basePath+"/testcase/test_%v", p)
			}
			fmt.Println(destPath)
			runTestCase.ChDirRun(basePath, destPath, 0)
		}
	}
	if len(args) == 0 {
		newargs := []string{"P0", "P1", "P2", "P3"}
		for _, p := range newargs {
			destPath := ""
			if runtime.GOOS == "windows" {
				destPath = fmt.Sprintf(basePath+"\\testcase\\test_%v", p)
				fmt.Println("Test123", destPath)
			} else {
				destPath = fmt.Sprintf(basePath+"/testcase/test_%v", p)
			}
			fmt.Println(destPath)
			runTestCase.ChDirRun(basePath, destPath, 0)
		}
	}
}

func main() {
	runLevel()
	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)
}
