package cli

import (
	"github.com/urfave/cli"

	"github.com/taqboz/tombo_gdn/cli/target"
	"github.com/taqboz/tombo_gdn/cli/view"
)

const version = "0.1.2"

func New() *cli.App {
	app := cli.NewApp()

	app.Name = "tombo_gdn"
	app.Usage = "This app check website's error information."
	app.Version = version

	app.Action = func (context *cli.Context) error {
		if err := target.InputTargetInfo(context.Args().Get(0)); err != nil {
			return err
		}

		// 調査開始の宣言
		view.FromXml()
		// ホスト内ディレクトリの取得
		if err := target.GetPageDirectory(); err != nil {
			return err
		}
		// 取得した情報の表示
		view.InfoGot()

		switch {
		case context.Bool("meta"):
			target.Option = "head > "
			if err := process("contents"); err != nil {
				return err
			}

		case context.Bool("body"):
			target.Option = "body > "
			if err := process("contents"); err != nil {
				return err
			}

		default:
			if err := process("contents"); err != nil {
				return err
			}
		}

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag {
			Name: "",
			Usage: "check all link in the target website.",
		},
		cli.BoolFlag {
			Name: "meta m",
			Usage: "check head contents",
		},
		cli.BoolFlag {
			Name: "body b",
			Usage: "check body contents",
		},
	}

	return app
}
