package kustomize

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/deislabs/porter/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

type InstallTest struct {
	expectedCommand string
	installStep     InstallStep
}

// sad hack: not sure how to make a common test main for all my subpackages
func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

func TestMixin_UnmarshalInstallStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-robotshop-input.yaml")
	require.NoError(t, err)

	var action InstallAction
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	//	require.Len(t, action.Steps, 1)
	step := action.Steps[0]

	//	assert.Equal(t, "Install MySQL", step.Description)
	//	assert.NotEmpty(t, step.Outputs)
	//	assert.Equal(t, KustomizeOutput{"mysql-root-password", "porter-ci-mysql", "mysql-root-password"}, step.Outputs[0])

	assert.Equal(t, "porter-robotshop-cart", step.Name)
	//	assert.Equal(t, '{"kustomize/robotshop/overlays/local/cart"', step.Kustomization)
	//assert.Equal(t, "0.10.2", step.Version)
	//assert.Equal(t, true, step.Replace)
	//assert.Equal(t, map[string]string{"mysqlDatabase": "mydb", "mysqlUser": "myuser",
	//	"livenessProbe.initialDelaySeconds": "30", "persistence.enabled": "true"}, step.Set)
}

func TestMixin_Install(t *testing.T) {
	name := "MYRELEASE"
	kustomization := []string{"MYKUSTOMIZATION"}
	reorder := "legacy"

	baseInstall := fmt.Sprintf(`kustomize build --name %s %s`, name, kustomization)

	installTests := []InstallTest{
		{
			expectedCommand: fmt.Sprintf(`%s`, baseInstall),
			installStep: InstallStep{
				InstallArguments: InstallArguments{
					Step:          Step{Description: "Install Robotshop"},
					Name:          name,
					Kustomization: kustomization,
					Reorder:       reorder,
				},
			},
		},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, installTest := range installTests {
		t.Run(installTest.expectedCommand, func(t *testing.T) {
			err := os.Setenv(test.ExpectedCommandEnv, installTest.expectedCommand)

			if err != nil {
				os.Exit(-1)
			}
			action := InstallAction{Steps: []InstallStep{installTest.installStep}}
			b, _ := yaml.Marshal(action)

			h := NewTestMixin(t)
			h.In = bytes.NewReader(b)

			err = h.Install()

			require.NoError(t, err)
		})
	}
}
