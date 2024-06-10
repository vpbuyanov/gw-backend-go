package configs

import (
	"fmt"
	"net/url"
)

type Kafka struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Password string `yaml:"password"`
}

func (p *Kafka) String() string {
	u := url.URL{
		Scheme: "kafka",
		User:   url.UserPassword(p.User, p.Password),
		Host:   fmt.Sprintf("%s:%d", p.Host, p.Port),
	}

	q := u.Query()
	q.Set("mechanism", "PLAIN")
	q.Set("protocol", "SASL_PLAINTEXT")

	u.RawQuery = q.Encode()

	return u.String()
}
