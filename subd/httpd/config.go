package httpd

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/megamsys/libgo/cmd"
)

type Config struct {
	Enabled     bool   `toml:"enabled"`
	BindAddress string `toml:"bind_address"`
	UseTls      bool   `toml:"use_tls"`
	CertFile    string `toml:"cert_file"`
	KeyFile     string `toml:"key_file"`
}

func (c Config) String() string {
	w := new(tabwriter.Writer)
	var b bytes.Buffer
	w.Init(&b, 0, 8, 0, '\t', 0)
	b.Write([]byte(cmd.Colorfy("\nConfig:", "white", "", "bold") + "\t" +
		cmd.Colorfy("httpd", "cyan", "", "") + "\n"))
	b.Write([]byte("enabled     " + "\t" + strconv.FormatBool(c.Enabled) + "\n"))
	b.Write([]byte("bind_address" + "\t" + c.BindAddress + "\n"))
	b.Write([]byte("usetls      " + "\t" + strconv.FormatBool(c.UseTls) + "\n"))
	b.Write([]byte("---\n"))
	fmt.Fprintln(w)
	w.Flush()
	return strings.TrimSpace(b.String())
}

func NewConfig() *Config {
	return &Config{
		Enabled:     true,
		BindAddress: "localhost:9005",
		UseTls:      false,
	}
}
