package caddy // import "go.zoe.im/gopkg/caddy"

import "github.com/mholt/caddy"

func init() {
	caddy.RegisterPlugin("gopkg", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	return nil
}
