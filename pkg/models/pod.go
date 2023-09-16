package models

type Pod struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Containers Containers `yaml:"containers"`
}

type Ports []struct {
	ContainerPort int `yaml:"containerPort"`
}

type Containers []struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Ports Ports  `yaml:"ports"`
}
