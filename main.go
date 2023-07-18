package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Cloud provider argument is missing")
	}

	cloudProvider := os.Args[1]
	err := os.Chdir(cloudProvider)
	if err != nil {
		log.Fatalf("Failed to change directory to %s: %v", cloudProvider, err)
	}

	// Run the Go program in the selected cloud provider folder
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
