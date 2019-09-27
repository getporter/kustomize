package kustomize

import (
	"strings"
	"testing"

	"github.com/deislabs/porter/pkg/printer"

	"github.com/deislabs/porter/pkg/porter/version"
	"github.com/stretchr/testify/require"

	"github.com/donmstewart/porter-kustomize/pkg"
)

func TestPrintVersion(t *testing.T) {
	pkg.Commit = "abc123"
	pkg.Version = "v1.2.3"

	m := NewTestMixin(t)

	opts := version.Options{}
	err := opts.Validate()
	require.NoError(t, err)
	m.PrintVersion(opts)

	gotOutput := m.TestContext.GetOutput()
	wantOutput := "kustomize v0.1 (alpha) by donmstewart"
	if !strings.Contains(gotOutput, wantOutput) {
		t.Fatalf("invalid output:\nWANT:\t%q\nGOT:\t%q\n", wantOutput, gotOutput)
	}
}

func TestPrintJsonVersion(t *testing.T) {
	pkg.Commit = "alpha"
	pkg.Version = "v0.1"

	m := NewTestMixin(t)

	opts := version.Options{}
	opts.RawFormat = string(printer.FormatJson)
	err := opts.Validate()
	require.NoError(t, err)
	m.PrintVersion(opts)

	gotOutput := m.TestContext.GetOutput()
	wantOutput := `{
	  "name": "kustomize",
	  "version": "v0.1",
	  "commit": "alpha",
	  "author": "donmstewart"
	}`

	if !strings.Contains(gotOutput, wantOutput) {
		t.Fatalf("invalid output:\nWANT:\t%q\nGOT:\t%q\n", wantOutput, gotOutput)
	}
}