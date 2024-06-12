package configs

type JWT struct {
	SecretAuth  string `yaml:"secret"`
	SecretReset string `yaml:"secret_reset"`
}
