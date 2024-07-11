package configs

type Mailer struct {
	Name              string `yaml:"name"`
	FromEmailAddress  string `yaml:"from_email_address"`
	FromEmailPassword string `yaml:"from_email_password"`
}
