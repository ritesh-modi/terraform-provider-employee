package main
 
import (
      "github.com/hashicorp/terraform-plugin-sdk/plugin"
      "github.com/hashicorp/terraform-plugin-sdk/terraform"
      "github.com/ritesh-modi/terraform-provider-employee/employee"
)
 
func main() {
      plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
                return employee.Provider()
        },
      })
}
