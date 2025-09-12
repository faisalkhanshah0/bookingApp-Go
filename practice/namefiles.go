package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	env := os.Getenv("JAVA_HOME")
	fmt.Println("value of JAVA_HOME is ", env)
	cmd := exec.Command("ls")
	output, _ := cmd.Output()
	lines := strings.Split(string(output), "\n")
	for i, k := range lines {
		fmt.Println(i, " ==== ", k)
	}
}
