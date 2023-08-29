// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/terramate-io/tf/grpcwrap"
	"github.com/terramate-io/tf/plugin"
	simple "github.com/terramate-io/tf/provider-simple"
	"github.com/terramate-io/tf/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
