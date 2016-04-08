package httpd

import (
	"testing"

	"github.com/BurntSushi/toml"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct {
	service *Service
}

var _ = check.Suite(&S{})

// Ensure the configuration can be parsed.
func (s *S) TestHttpdConfig_Parse(c *check.C) {
	// Parse configuration.
	var cm Config
	if _, err := toml.Decode(`
enabled = true
bind_address = ":8080"
use_tls =  false
`, &cm); err != nil {
		c.Fatal(err)
	}

	c.Assert(cm.BindAddress, check.Equals, ":8080")
	c.Assert(cm.UseTls, check.Equals, false)
}
