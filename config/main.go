package config

type Integration struct {
	Namespace string   `json:"namespace"`
	NoAck     []string `json:"noack"`
}

type BufIoConfig struct {
	Config struct {
		Providers []string `json:"providers"`
		Integrations struct {
			Slack Integration `json:"slack"`
		}
	} `json:"config"`
}