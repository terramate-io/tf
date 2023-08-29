// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clistate

import (
	"testing"

	"github.com/terramate-io/tf/command/arguments"
	"github.com/terramate-io/tf/command/views"
	"github.com/terramate-io/tf/states/statemgr"
	"github.com/terramate-io/tf/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
