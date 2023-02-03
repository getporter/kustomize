package kustomize

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// The `Porter.sh` action for Install
type InstallAction struct {
	Steps []InstallStep `yaml:"install"`
}

// The `Porter.sh` step for Install for Kustomize
type InstallStep struct {
	InstallArguments `yaml:"kustomize"`
}

// The base level Structure that captures the high level data types
//
//	needed by Kustomize.
//
//	`Kustomization` field in the Go struct and `kustomization_input` field in the `porter.yaml`
//	is the location against which to run the `kustomize build` command. This will
//	be an overlay directory.
//
//	`Manifests` field in the Go struct and `kubernetes_manifest_output` field in the `porter.yaml`
//	is the location into which `kustomize` will output the generated kubernetes resource yaml files.
//
//	`Reorder` is a boolean flag in the Go struct and `autoDeploy` field in the `porter.yaml` which
//	 enables `kustomize` to reorder the resources within the yaml file that is to be output.
//
//	 `Reorder` from the `kustomize` documentation -
//
//	 --reorder {none | legacy } flag to the build command.
//	 The default value is legacy which means no change - continue to output resources in the legacy order
//	 (Namespaces first, ValidatingWebhookConfiguration last, etc. - see gvk.go)
//	 A value of none suppresses the sort
//
//	`Set`, `AutoDeploy` are not currently implemented.
type InstallArguments struct {
	Step `yaml:",inline"`

	Name          string            `yaml:"name"`
	Kustomization []string          `yaml:"kustomization_input"`
	Manifests     string            `yaml:"kubernetes_manifest_output"`
	Set           map[string]string `yaml:"set"`
	AutoDeploy    bool              `yaml:"autoDeploy"`
	Reorder       string            `yaml:"reorder"`
}

// The public method invoked by `porter` when performing an `Install` step that has a `kustomize` mixin step
func (m *Mixin) Install(ctx context.Context) error {
	payload, err := m.getPayloadData()
	if err != nil {
		return err
	}

	var action InstallAction
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		return err
	}
	if len(action.Steps) != 1 {
		return errors.Errorf("expected a single step, but got %d", len(action.Steps))
	}

	step := action.Steps[0]

	var commands []*exec.Cmd

	ghToken := step.Set["kustomizeBaseGHToken"]

	err = m.configureGithubToken(ctx, ghToken)
	if err != nil {
		return err
	}

	err = m.manifestHandling(step)
	if err != nil {
		return err
	}

	err = m.buildAndExecuteKustomizeCmds(ctx, step, commands)
	if err != nil {
		return err
	}

	for _, output := range step.Outputs {
		err = m.Context.WriteMixinOutputToFile(output.Name, []byte(fmt.Sprintf("%v", output)))
		if err != nil {
			return errors.Wrapf(err, "unable to write output '%s'", output.Name)
		}
	}
	return nil
}
