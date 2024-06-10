package configs

import "fmt"

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
