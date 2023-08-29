package moduletest

import (
	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/tfdiags"
)

type File struct {
	Config *configs.TestFile

	Name   string
	Status Status

	Runs []*Run

	Diagnostics tfdiags.Diagnostics
}
