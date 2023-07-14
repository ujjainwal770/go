package main
import "fmt"
import (
	"context"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Get the command-line arguments
	apiSecret := os.Args[2]
	apiKey := os.Args[4]
	smtpUser := os.Args[6]
	smtpPassword := os.Args[8]
	ceebitLicenseKey := os.Args[10]
	ceebitLicenseSecret := os.Args[12]
	ceebitCodeUser := os.Args[14]
	ceebitCodePassword := os.Args[16]

	// Set environment variables for Terraform
	os.Setenv("GODADDY_API_SECRET", apiSecret)
	os.Setenv("GODADDY_API_KEY", apiKey)
	os.Setenv("SMTP_USER", smtpUser)
	os.Setenv("SMTP_PASSWORD", smtpPassword)
	os.Setenv("CEEBIT_LICENSE_KEY", ceebitLicenseKey)
	os.Setenv("CEEBIT_LICENSE_SECRET", ceebitLicenseSecret)
	os.Setenv("CEEBIT_CODE_USER", ceebitCodeUser)
	os.Setenv("CEEBIT_CODE_PASSWORD", ceebitCodePassword)

	// Run Terraform commands
	initCmd := exec.Command("terraform", "init")
	err := initCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	applyCmd := exec.Command("terraform", "apply", "-auto-approve")
	err = applyCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Terraform apply completed successfully.")
}
