package main

import (
	"github.com/donmstewart/porter-kustomize/pkg/kustomize"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *kustomize.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the upgrade functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Upgrade(cmd.Context())
		},
	}
	return cmd
}
