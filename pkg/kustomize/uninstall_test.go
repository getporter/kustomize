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

type UninstallTest struct {
	expectedCommand string
	uninstallStep   UninstallStep
}

func TestMixin_UnmarshalUninstallStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/uninstall-input.yaml")
	require.NoError(t, err)

	var action UninstallAction
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)
	step := action.Steps[0]

	assert.Equal(t, "Uninstall MySQL", step.Description)
	assert.True(t, step.Purge)
}

func TestMixin_Uninstall(t *testing.T) {

	uninstallTests := []UninstallTest{
		{
			expectedCommand: `kustomize delete foo bar`,
			uninstallStep: UninstallStep{
				UninstallArguments: UninstallArguments{
					Step: Step{Description: "Uninstall Foo"},
				},
			},
		},
		{
			expectedCommand: `kustomize delete --purge foo bar`,
			uninstallStep: UninstallStep{
				UninstallArguments: UninstallArguments{
					Step:  Step{Description: "Uninstall Foo"},
					Purge: true,
				},
			},
		},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, uninstallTest := range uninstallTests {
		t.Run(uninstallTest.expectedCommand, func(t *testing.T) {
			err := os.Setenv(test.ExpectedCommandEnv, uninstallTest.expectedCommand)
			if err != nil {
				os.Exit(-1)
			}

			action := UninstallAction{Steps: []UninstallStep{uninstallTest.uninstallStep}}
			b, _ := yaml.Marshal(action)

			x := string(b)
			fmt.Println(x)
			h := NewTestMixin(t)
			h.In = bytes.NewReader(b)

			err = h.Uninstall()

			require.NoError(t, err)
		})
	}
}
