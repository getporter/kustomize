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

type UpgradeTest struct {
	expectedCommand string
	upgradeStep     UpgradeStep
}

func TestMixin_UnmarshalUpgradeStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/upgrade-input.yaml")
	require.NoError(t, err)

	var action UpgradeAction
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)
	step := action.Steps[0]

	assert.Equal(t, "Upgrade MySQL", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, KustomizeOutput{"mysql-root-password", "porter-ci-mysql", "mysql-root-password"}, step.Outputs[0])

	assert.Equal(t, "porter-ci-mysql", step.Name)
	assert.Equal(t, "stable/mysql", step.Kustomization)
	assert.Equal(t, "0.10.2", step.Version)
	assert.True(t, step.Wait)
	assert.True(t, step.ResetValues)
	assert.True(t, step.ResetValues)
	assert.Equal(t, map[string]string{"mysqlDatabase": "mydb", "mysqlUser": "myuser",
		"livenessProbe.initialDelaySeconds": "30", "persistence.enabled": "true"}, step.Set)
}

func TestMixin_Upgrade(t *testing.T) {
	name := "MYRELEASE"
	kustomization := "MYKUSTOMIZATION"

	baseUpgrade := fmt.Sprintf(`kustomize upgrade %s %s`, name, kustomization)

	upgradeTests := []UpgradeTest{
		{
			expectedCommand: baseUpgrade,
			upgradeStep: UpgradeStep{
				UpgradeArguments: UpgradeArguments{
					Step:          Step{Description: "Upgrade Foo"},
					Name:          name,
					Kustomization: kustomization,
				},
			},
		},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, upgradeTest := range upgradeTests {
		t.Run(upgradeTest.expectedCommand, func(t *testing.T) {

			err := os.Setenv(test.ExpectedCommandEnv, upgradeTest.expectedCommand)
			if err != nil {
				os.Exit(-1)
			}

			action := UpgradeAction{Steps: []UpgradeStep{upgradeTest.upgradeStep}}
			b, err := yaml.Marshal(action)
			require.NoError(t, err)

			h := NewTestMixin(t)
			h.In = bytes.NewReader(b)

			err = h.Upgrade()

			require.NoError(t, err)
		})
	}
}
