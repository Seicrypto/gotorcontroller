package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "torcontroller -V")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}
