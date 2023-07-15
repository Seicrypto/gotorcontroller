package gocontroller

import (
	"fmt"
	"os/exec"
)

func torStart() {
	cmd := exec.Command("bash", "-c", "torcontroller -S")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}

func torStop() {
	cmd := exec.Command("bash", "-c", "torcontroller -C")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}

func torSwitch() {
	cmd := exec.Command("bash", "-c", "torcontroller -W")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}
