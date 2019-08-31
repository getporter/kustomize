package kustomize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	m := NewTestMixin(t)

	err := m.Build()
	require.NoError(t, err)

	wantOutput := `RUN apt-get update && \
 apt-get install -y curl git && \
 curl -L -O https://github.com/kubernetes-sigs/kustomize/releases/download/v3.1.0/kustomize_3.1.0_linux_amd64 && \
 mv ./kustomize_3.1.0_linux_amd64 /usr/local/bin/kustomize && \
 chmod a+x /usr/local/bin/kustomize
`

	gotOutput := m.TestContext.GetOutput()
	assert.Equal(t, wantOutput, gotOutput)
}
