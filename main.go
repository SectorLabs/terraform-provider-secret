package main

import (
	"github.com/SectorLabs/terraform-provider-secret/secret"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: secret.Provider})
}
