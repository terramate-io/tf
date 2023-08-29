// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/terramate-io/tf/builtin/providers/terraform"
	"github.com/terramate-io/tf/grpcwrap"
	"github.com/terramate-io/tf/plugin"
	"github.com/terramate-io/tf/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
