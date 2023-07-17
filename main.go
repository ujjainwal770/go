package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cloudProvider := os.Args[1]

	switch cloudProvider {
	case "AWS":
		// Set up AWS-specific configurations and execute Terraform commands
		fmt.Println("Configuring AWS provider...")
		runTerraformCommand("init", "aws")
		runTerraformCommand("plan", "aws")
		runTerraformCommand("apply", "aws")
	case "Azure":
		// Set up Azure-specific configurations and execute Terraform commands
		fmt.Println("Configuring Azure provider...")
		runTerraformCommand("init", "azure")
		runTerraformCommand("plan", "azure")
		runTerraformCommand("apply", "azure")
	case "Google Cloud":
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
	cmd := exec.Command("terraform", command, "-var-file=terraform.tfvars", "-var-file="+provider+".tfvars")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
