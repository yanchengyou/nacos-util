package model

import "fmt"

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func (c Config) String() string {
	return fmt.Sprintf("username: %s\npassword: %s\nhost: %s", c.Username, c.Password, c.Host)
}

