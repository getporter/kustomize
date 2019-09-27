package kustomize

import (
	"fmt"
)

// Outputs the JSON Schema for this `kustomize` mixin's configuration found within a Porter.sh
//   `porter.yaml` configuration file.
func (m *Mixin) PrintSchema() error {
	schema, err := m.GetSchema()
	if err != nil {
		return err
	}

	fmt.Fprintf(m.Out, schema)

	return nil
}

// Returns the JSON Schema for this `kustomize` mixin's configuration, found within a Porter.sh
//   `porter.yaml` configuration file, as a `string`
func (m *Mixin) GetSchema() (string, error) {
	b, err := m.schema.Find("kustomize.json")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
