package kustomize

import (
	"fmt"
)

const kustomizeClientVersion = "v2.12.3"
const dockerfileLines = `RUN apt-get update && \
 apt-get install -y curl && \
 curl -o kustomize.tgz https://storage.googleapis.com/kubernetes-kustomize/kustomize-%s-linux-amd64.tar.gz && \
 tar -xzf kustomize.tgz && \
 mv linux-amd64/kustomize /usr/local/bin && \
 rm kustomize.tgz
RUN kustomize init --client-only`

func (m *Mixin) Build() error {
	fmt.Fprintf(m.Out, dockerfileLines, kustomizeClientVersion)
	return nil
}
