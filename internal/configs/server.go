package configs

import "fmt"

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
