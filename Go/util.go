package main

import (
	// "fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetInput(filepath string) string {
	prefix := "../inputs/"
	dat, err := os.ReadFile(prefix + filepath)
	if err != nil {
		panic(err)
	}

	inputString := string(dat[:])
	inputString = strings.TrimSpace(inputString)
	return inputString
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
