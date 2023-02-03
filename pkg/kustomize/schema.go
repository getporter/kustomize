package kustomize

import (
	"fmt"
)

// Outputs the JSON Schema for this `kustomize` mixin's configuration found within a Porter.sh
//
//	`porter.yaml` configuration file.
func (m *Mixin) PrintSchema() error {
	fmt.Fprintf(m.Out, schema)

	return nil
}
