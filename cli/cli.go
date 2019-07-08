package cli

import "github.com/urfave/cli"

const version = "0.2.0"

func New() *cli.App {
	app := cli.NewApp()

	app.Name = "tombo_gdn"
	app.Usage = "This app check website's error information."

	app.Version = version
	app.Author = "taqboz"

	app.Commands = []cli.Command{
		{
			Name: "check",
			Aliases: []string{"c"},
			Usage: "check all contents in website",
			Action: func(c *cli.Context) error {
				if err := check(c.Args().First()); err != nil {
					return err
				}
				return nil
			},
		},
	}

	return app
}
