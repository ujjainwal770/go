provider "azurerm" {
  version = ">= 2.0.0"
  features {}
}

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
    }
  }
}
