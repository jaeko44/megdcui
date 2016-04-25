package meta

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"strings"
	"text/tabwriter"
    "path/filepath"
	"github.com/megamsys/libgo/cmd"
)

const (
	// DefaultScylla is the default scylla if one is not provided.
	DefaultScylla = "192.168.1.100"

	// DefaultScyllaKeyspace is the default Scyllakeyspace if one is not provided.
	DefaultScyllaKeyspace = "vertice"

	MEGAM_HOME = "MEGAM_HOME"
	// DefaultNSQ is the default nsqd if its not provided.
	DefaultNSQd = "localhost:4161"
)

// Config represents the meta configuration.
type Config struct {
	Dir            string   `toml:"dir"`
	Scylla         []string `toml:"scylla"`
	ScyllaKeyspace string   `toml:"scylla_keyspace"`
	NSQd []string `toml:"nsqd"`
	Api 		string `toml:"api"`
}

var MC *Config

func (c Config) String() string {
	w := new(tabwriter.Writer)
	var b bytes.Buffer
	w.Init(&b, 0, 8, 0, '\t', 0)
	b.Write([]byte(cmd.Colorfy("Config:", "white", "", "bold") + "\t" +
		cmd.Colorfy("Meta", "cyan", "", "") + "\n"))
	b.Write([]byte("Dir       " + "\t" + c.Dir + "\n"))
	b.Write([]byte("Scylla    " + "\t" + strings.Join(c.Scylla, ",") + "\n"))
	b.Write([]byte("ScyllaKeyspace" + "\t" + c.ScyllaKeyspace + "\n"))
	b.Write([]byte("---\n"))
	fmt.Fprintln(w)
	w.Flush()
	return strings.TrimSpace(b.String())
}

func NewConfig() *Config {
	var homeDir string
	// By default, store logs, meta and load conf files in MEGAM_HOME directory
	if os.Getenv(MEGAM_HOME) != "" {
		homeDir = os.Getenv(MEGAM_HOME)
	} else if u, err := user.Current(); err == nil {
		homeDir = u.HomeDir
	} else {
		return nil
	}

	defaultDir := filepath.Join(homeDir, "megdc/")

	// Config represents the configuration format for the vertice.
	return &Config{
		Dir:            defaultDir,
		Scylla:         []string{DefaultScylla},
		ScyllaKeyspace: DefaultScyllaKeyspace,
		NSQd: []string{DefaultNSQd},
	}
}

func (c *Config) ToMap() map[string]string {
	mp := make(map[string]string)
	mp["dir"] = c.Dir
	mp["scylla_host"] = strings.Join(c.Scylla, ",")
	mp["scylla_keyspace"] = c.ScyllaKeyspace
	return mp
}

func (c *Config) MkGlobal() {
	MC = c
}
