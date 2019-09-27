package kustomize

import (
	"github.com/deislabs/porter/pkg/mixin"
	"github.com/deislabs/porter/pkg/porter/version"
	"github.com/donmstewart/porter-kustomize/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "kustomize",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "donmstewart",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
