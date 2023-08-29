// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/terramate-io/tf/grpcwrap"
	plugin "github.com/terramate-io/tf/plugin6"
	simple "github.com/terramate-io/tf/provider-simple-v6"
	"github.com/terramate-io/tf/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
