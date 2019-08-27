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
 apt-get install -y curl && \
 curl -O https://github.com/kubernetes-sigs/kustomize/releases/download/v3.1.0/kustomize_3.1.0_linux_amd64 && \
 mv ./kustomize_3.1.0_linux_amd64 /usr/local/bin/kustomize && \
 chmod a+x /usr/local/bin/kustomize
RUN apt-get update && \
 apt-get install -y apt-transport-https curl && \
 curl -o kubectl https://storage.googleapis.com/kubernetes-release/release/v1.15.3/bin/linux/amd64/kubectl && \
 mv kubectl /usr/local/bin && \
 chmod a+x /usr/local/bin/kubectl`

	gotOutput := m.TestContext.GetOutput()
	assert.Equal(t, wantOutput, gotOutput)
}
