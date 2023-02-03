package main

import (
	"github.com/donmstewart/porter-kustomize/pkg/kustomize"
	"github.com/spf13/cobra"
)

func buildStatusCommand(m *kustomize.Mixin) *cobra.Command {
	opts := kustomize.StatusOptions{}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Print the status of the kustomize components in the bundle",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.ParseFormat()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Status(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVarP(&opts.RawFormat, "output", "o", "plaintext", "Output format. Allowed values are: plaintext, yaml, json")
	return cmd
}
