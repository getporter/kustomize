package kustomize

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os/exec"
)

type InstallAction struct {
	Steps []InstallStep `yaml:"install"`
}

type InstallStep struct {
	InstallArguments `yaml:"kustomize"`
}

type InstallArguments struct {
	Step `yaml:",inline"`

	Name          string            `yaml:"name"`
	Kustomization []string          `yaml:"kustomization_input"`
	Manifests     string            `yaml:"kubernetes_manifest_output"`
	Set           map[string]string `yaml:"set"`
	AutoDeploy    bool              `yaml:"autoDeploy"`
	Reorder       string            `yaml:"reorder"`
}

func (m *Mixin) Install() error {
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

	err = m.configureGithubToken(ghToken)
	if err != nil {
		return err
	}

	err = m.manifestHandling(step)
	if err != nil {
		return err
	}

	err = m.buildAndExecuteKustomizeCmds(step, commands)
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
