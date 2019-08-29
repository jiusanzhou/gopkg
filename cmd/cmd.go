package cmd

import (
	"log"

	"go.zoe.im/x/cli"
)

var (
	// global cmd contains all sub command
	cmd = cli.New(
		cli.Name("gopkg"),
		cli.Short("Gopkg allows you to create vanity Go import urls."),
		cli.Run(func(c *cli.Command, args ...string) {
			c.Help()
		}),
	)
)

func init() {
	// Register other commands

	svr := newServer()
	gen := newGenerator()

	cmd.Register(
		cli.New(
			cli.Name("serve"),
			cli.Short("Start a http server."),
			cli.SetFlags(func(c *cli.Command) {
				c.Flags().StringVarP(&svr.addr, "addr", "", ":8080", "address to listen.")
				c.Flags().StringVarP(&svr.index, "index", "", "https://zoe.im", "home page.")
			}),
			cli.Run(func(c *cli.Command, args ...string) {
				err := svr.run()
				if err != nil {
					log.Println("[gopkg] [server]", err)
				}
			}),
		),
		cli.New(
			cli.Name("gen"),
			cli.Short("Generate htmls for static provider."),
			cli.SetFlags(func(c *cli.Command) {
				c.Flags().StringVarP(&gen.target, "output", "o", "docs", "Output of generated files.")
				c.Flags().StringVarP(&gen.configPath, "config", "c", "gopkg.yaml", "Configuration file to generate pkg.")
				c.Flags().BoolVarP(&gen.debug, "debug", "d", false, "Debug.")
			}),
			cli.Run(func(c *cli.Command, args ...string) {
				err := gen.run()
				if err != nil {
					log.Println("[gopkg] [gen]", err)
				}
			}),
		),
	)
}

// Register create a sub command
func Register(cmds ...*cli.Command) error {
	return cmd.Register(cmds...)
}

// Run execute command
func Run(opts ...cli.Option) error {
	return cmd.Run(opts...)
}
