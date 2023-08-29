// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import "github.com/terramate-io/tf/command/workdir"

func WorkingDir(originalDir string, overrideDataDir string) *workdir.Dir {
	ret := workdir.NewDir(".") // caller should already have used os.Chdir in "-chdir=..." mode
	ret.OverrideOriginalWorkingDir(originalDir)
	if overrideDataDir != "" {
		ret.OverrideDataDir(overrideDataDir)
	}
	return ret
}
