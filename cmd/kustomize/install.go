package main

import (
	"github.com/donmstewart/porter-kustomize/pkg/kustomize"
	"github.com/spf13/cobra"
)

var (
	commandFile string
)

func buildInstallCommand(m *kustomize.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Install()
		},
	}
	return cmd
}
