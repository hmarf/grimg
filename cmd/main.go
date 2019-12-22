package main

import (
	"os"

	"github.com/hmarf/gregif/rimg"
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "grimg"
	app.Usage = "Compress images(jpg, png, gif)"
	app.Version = "0.0.1"
	app.Author = "hmarf"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Value: "None",
			Usage: "[required] string\n	input file path",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "output",
			Usage: "string\n	output file path [default] output.(file type)",
		},
		cli.Float64Flag{
			Name:  "compression, c",
			Value: 0.0,
			Usage: "[required] float(0.1 ~ 0.9)\n	Specify compression ratio",
		},
	}
	return app
}

func Action(c *cli.Context) {
	app := App()
	if c.String("input") == "None" {
		app.Run(os.Args)
		return
	}
	if c.Float64("compression") == 0.0 {
		app.Run(os.Args)
		return
	}
	option := rimg.Option{
		InputFile:   c.String("input"),
		OutputFile:  c.String("output"),
		Compression: c.Float64("compression"),
	}
	rimg.Grimg(option)
}

func main() {
	app := App()
	app.Action = Action
	app.Run(os.Args)
}
