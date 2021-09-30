package model

import "fmt"

type Config struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Config) String() string {
	return fmt.Sprintf("username: %s\npassword: %s\nhost: %s", c.Username, c.Password, c.Host)
}

