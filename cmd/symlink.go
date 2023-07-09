package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	basePath, _ := os.Getwd()
	// Source symlink directory
	source := filepath.Join(basePath + "/storage/app/public/")
	// Target directory
	destination := filepath.Join(basePath + "/public/uploads/")
	// Create symlink
	cmd := exec.Command("ln", "-s", source, destination)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error symlink:", err)
		return
	}

	fmt.Println("Successfully create symlink!")
}
