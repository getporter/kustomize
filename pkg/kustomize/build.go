package kustomize

import (
	"bytes"
	"fmt"
	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/Masterminds/semver"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"text/template"
)

const (
	dockerfileLines string = `RUN apt-get update && \
apt-get install -y curl git && \
curl -L -O https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F{{ .KustomizeClientVersion }}/kustomize_{{ .KustomizeClientVersion }}_linux_amd64.tar.gz && \
tar -zxvf kustomize_{{ .KustomizeClientVersion }}_linux_amd64.tar.gz && \
mv ./kustomize /usr/local/bin/kustomize && \
chmod a+x /usr/local/bin/kustomize
`
	// clientVersionConstraint represents the semver constraint for the Kustomize client version
	// Currently, this mixin only supports Kustomize clients versioned v3.x.x
	clientVersionConstraint string = "^v3.x"
)

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

func (m *Mixin) Build() error {
	// Create new Builder.
	var input BuildInput
	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "" {
		ok, err := validate(suppliedClientVersion, clientVersionConstraint)
		if err != nil {
			return err
		}
		if !ok {
			return errors.Errorf("supplied clientVersion %q does not meet semver constraint %q",
				suppliedClientVersion, clientVersionConstraint)
		}
		m.KustomizeClientVersion = suppliedClientVersion
	}

	t, err := template.New("cmd").Parse(dockerfileLines)
	if err != nil {
		return err
	}

	var cmd bytes.Buffer
	err = t.Execute(&cmd, m)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(m.Out, cmd.String())
	if err != nil {
		return err
	}
	return nil
}

// validate validates that the supplied clientVersion meets the supplied semver constraint
func validate(clientVersion, constraint string) (bool, error) {
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false, errors.Wrapf(err, "unable to parse version constraint %q", constraint)
	}

	v, err := semver.NewVersion(clientVersion)
	if err != nil {
		return false, errors.Wrapf(err, "supplied client version %q cannot be parsed as semver", clientVersion)
	}

	return c.Check(v), nil
}
