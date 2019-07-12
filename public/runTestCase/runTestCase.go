package runTestCase

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ChDirRun(basePath, destPath string, tag int) string {
	//timestamp := time.Now().Unix()
	//tm := time.Unix(timestamp, 0).Format("2006-01-02_03-04-05")
	defer func() {
		if err := recover(); err != nil {
			_ = os.Chdir(basePath)
			fmt.Println("+")
		}
	}()
	err := os.Chdir(destPath)
	if err != nil {
		log.Fatal("Failed to change dir")
	}
	fmt.Println(os.Getwd())
	//cmd := exec.Command("ginkgo", "-r")
	cmd := exec.Command("go", "test", "./...")
	fmt.Println(cmd)
	if tag == 1 {
		cmd = exec.Command("go", "test", "./...")
	}
	err = cmd.Run()
	if err != nil {
		log.Println("Run command error:", err)
		panic("Run command error")
	}
	return ""
}
