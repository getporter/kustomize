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
 curl -o kustomize.tgz https://storage.googleapis.com/kubernetes-kustomize/kustomize-v2.12.3-linux-amd64.tar.gz && \
 tar -xzf kustomize.tgz && \
 mv linux-amd64/kustomize /usr/local/bin && \
 rm kustomize.tgz
RUN kustomize init --client-only`

	gotOutput := m.TestContext.GetOutput()
	assert.Equal(t, wantOutput, gotOutput)
}
