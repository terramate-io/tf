// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/terramate-io/tf/command/jsonformat/computed"
	"github.com/terramate-io/tf/command/jsonformat/computed/renderers"
	"github.com/terramate-io/tf/command/jsonformat/structured"
)

func computeAttributeDiffAsPrimitive(change structured.Change, ctype cty.Type) computed.Diff {
	return asDiff(change, renderers.Primitive(change.Before, change.After, ctype))
}
