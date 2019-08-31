package kustomize

import "fmt"

const kustomizeClientVersion = "3.1.0"
const dockerfileLines string = `RUN apt-get update && \
 apt-get install -y curl git && \
 curl -L -O https://github.com/kubernetes-sigs/kustomize/releases/download/v%s/kustomize_%s_linux_amd64 && \
 mv ./kustomize_%s_linux_amd64 /usr/local/bin/kustomize && \
 chmod a+x /usr/local/bin/kustomize
`

func (m *Mixin) Build() error {
	var cmd = fmt.Sprintf(dockerfileLines, kustomizeClientVersion, kustomizeClientVersion, kustomizeClientVersion)
	_, err := fmt.Fprint(m.Out, cmd)
	if err != nil {
		return err
	}
	return nil
}
