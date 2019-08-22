package kustomize

type Step struct {
	Description string       `yaml:"description"`
	Outputs     []KustomizeOutput `yaml:"outputs,omitempty"`
}

type KustomizeOutput struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
	Key    string `yaml:"key"`
}
