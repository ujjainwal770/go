provider "aws" {
  region = "us-east-1" # Replace with your desired region
}

resource "aws_instance" "test_vm" {
  ami           = "ami-0c94855ba95c71c99" # Replace with your desired AMI ID
  instance_type = "t2.micro"               # Replace with your desired instance type

  tags = {
    Name = "TestVM"
  }
}