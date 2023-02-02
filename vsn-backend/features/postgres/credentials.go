package postgres

import (
	"fmt"
	"net/url"
)

var DefaultCredentials Credentials = Credentials{
	User:     "user",
	Password: "password",
	Database: "vsn-db",
	Address:  "localhost:5432",
	Params: url.Values{
		"sslmode": []string{"disable"},
	},
}

type Credentials struct {
	User     string
	Password string
	Address  string // "localhost:5432"
	Database string // "vsn-db"
	Params   url.Values
}

func (c *Credentials) String() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?%s",
		c.User, c.Password, c.Address, c.Database, c.Params.Encode())
}
