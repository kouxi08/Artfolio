package config

type Record struct {
	Name       string `json:"name"`
	RecordType string `json:"recordType"`
	TTL        string `json:"ttl"`
	Content    string `json:"content"`
}

type KubeConfig struct {
	DeploymentName string `json:"deploymentName"`
	ServiceName    string `json:"serviceName"`
	IngressName    string `json:"ingressName"`
	HostName       string `json:"hostName"`
}

type Config struct {
	Record     Record     `json:"record"`
	KubeConfig KubeConfig `json:"kubeconfig"`
}
