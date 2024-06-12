package configs

import (
	"fmt"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
}

func (p *Postgres) String() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(p.User, p.Password),
		Host:   fmt.Sprintf("%s:%d", p.Host, p.Port),
		Path:   p.DBName,
	}

	q := u.Query()
	q.Set("sslmode", "disable")

	u.RawQuery = q.Encode()

	return u.String()
}

func (p *Postgres) MigrationsUp(urls ...string) error {
	var sourceURL string
	if urls == nil {
		sourceURL = "file://migrations"
	} else {
		sourceURL = urls[0]
	}
	m, err := migrate.New(sourceURL, p.String())
	if err != nil {
		return fmt.Errorf("can not create New migrate, err: %w", err)
	}
	if err = m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("can not UP migrate, err: %w", err)
	}

	return nil
}
