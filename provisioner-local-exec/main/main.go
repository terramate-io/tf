// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	localexec "github.com/terramate-io/tf/builtin/provisioners/local-exec"
	"github.com/terramate-io/tf/grpcwrap"
	"github.com/terramate-io/tf/plugin"
	"github.com/terramate-io/tf/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
