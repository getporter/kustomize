package kustomize

import "fmt"

const kustomizeClientVersion = "3.1.0"
const dockerfileLines string = `RUN apt-get update && \
 apt-get install -y curl && \
 curl -O https://github.com/kubernetes-sigs/kustomize/releases/download/v%s/kustomize_%s_linux_amd64 && \
 mv ./kustomize_3.1.0_linux_amd64 /usr/local/bin/kustomize && \
 chmod a+x /usr/local/bin/kustomize
`

/*
// kubectl may be necessary; for example, to set up RBAC for Helm's Tiller component if needed
const kubeVersion string = "v1.15.3"
const getKubectl string = `RUN apt-get update && \
 apt-get install -y apt-transport-https curl && \
 curl -o kubectl https://storage.googleapis.com/kubernetes-release/release/%s/bin/linux/amd64/kubectl && \
 mv kubectl /usr/local/bin && \
 chmod a+x /usr/local/bin/kubectl
`

*/

func (m *Mixin) Build() error {
	fmt.Fprintf(m.Out, dockerfileLines, kustomizeClientVersion, kustomizeClientVersion)
	//fmt.Fprintf(m.Out, getKubectl, kubeVersion)
	return nil
}
