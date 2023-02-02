package postgres

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentials(t *testing.T) {
	credentials := Credentials{
		User:     "user",
		Password: "password",
		Database: "vsn-db",
		Address:  "localhost:5432",
		Params: url.Values{
			"sslmode": []string{"disable"},
		},
	}

	assert.Equal(t, credentials.String(), "postgresql://user:password@localhost:5432/vsn-db?sslmode=disable")
}
