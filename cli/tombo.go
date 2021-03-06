package cli

import "github.com/urfave/cli"

const (
	version = "0.7.0"
	author = "taqboz"
)

func New() *cli.App {
	app := cli.NewApp()

	app.Name = "tombo_gdn"
	app.Usage = "This app check website's error information."

	app.Version = version
	app.Author = author

	app.Commands = []cli.Command{
		{
			Name: "check",
			Aliases: []string{"c"},
			Usage: "すべての項目のチェック（リンクの有効性、SEO設定）",
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
