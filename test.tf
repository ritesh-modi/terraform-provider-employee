
terraform {
  required_providers {
    example = {
      source = "github.com/ritesh-modi/employee"
    }
  }
}
provider "example" {
    
}

resource "example_server" "my-server" {
  address = "1.2.3.4"
}