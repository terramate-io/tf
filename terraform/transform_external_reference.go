package terraform

import "github.com/terramate-io/tf/addrs"

// ExternalReferenceTransformer will add a GraphNodeReferencer into the graph
// that makes no changes to the graph itself but, by referencing the addresses
// within ExternalReferences, ensures that any temporary nodes that are required
// by an external caller, such as the terraform testing framework, are not
// skipped because they are not referenced from within the module.
type ExternalReferenceTransformer struct {
	ExternalReferences []*addrs.Reference
}

func (t *ExternalReferenceTransformer) Transform(g *Graph) error {
	if len(t.ExternalReferences) == 0 {
		return nil
	}

	g.Add(&nodeExternalReference{
		ExternalReferences: t.ExternalReferences,
	})
	return nil
}
