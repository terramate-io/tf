package moduletest

import (
	"github.com/terramate-io/tf/configs"
	"github.com/terramate-io/tf/tfdiags"
)

type File struct {
	Config *configs.TestFile

	Name   string
	Status Status

	Runs []*Run

	Diagnostics tfdiags.Diagnostics
}
