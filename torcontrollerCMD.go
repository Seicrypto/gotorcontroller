package gotorcontroller

// This go file build for calling torcontroller.deb bash command.
import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Return boolean to show bash command succeeded or not.
func Checktor() (bool, error) {
	cmd := exec.Command("bash", "-c", "systemctl status tor")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return strings.Contains(string(output), "running"), nil
}

func Checkprivoxy() (bool, error) {
	cmd := exec.Command("bash", "-c", "systemctl status privoxy")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return strings.Contains(string(output), "running"), nil
}

func StartTor() (bool, error) {
	cmd := exec.Command("bash", "-c", "torcontroller -S")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return true, nil
}

func StopTor() (bool, error) {
	cmd := exec.Command("bash", "-c", "torcontroller -C")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return true, nil
}

func SwitchIP() (bool, error) {
	cmd := exec.Command("bash", "-c", "torcontroller -W")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return true, nil
}
