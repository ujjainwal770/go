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

	switch cloudProvider {
	case "aws":
		// Set up AWS-specific configurations and execute Terraform commands
		fmt.Println("Configuring AWS provider...")
		runTerraformCommand("init", "aws")
		runTerraformCommand("plan", "aws")
		runTerraformCommand("apply", "aws")
	case "azure":
		// Set up Azure-specific configurations and execute Terraform commands
		fmt.Println("Configuring Azure provider...")
		runTerraformCommand("init", "azure")
		runTerraformCommand("plan", "azure")
		runTerraformCommand("apply", "azure")
	case "gcp":
		// Set up Google Cloud-specific configurations and execute Terraform commands
		fmt.Println("Configuring Google Cloud provider...")
		runTerraformCommand("init", "gcp")
		runTerraformCommand("plan", "gcp")
		runTerraformCommand("apply", "gcp")
	default:
		log.Fatal("Invalid cloud provider selected")
	}
}

func runTerraformCommand(command string, provider string) {
	cmd := exec.Command("terraform", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
