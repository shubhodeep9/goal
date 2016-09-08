package main

import (
	//"fmt"
	"github.com/gizak/termui"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "GOaL"
	app.Usage = "Go and Learn, A terminal based go learning tool"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "check",
			Value: "main.go",
			Usage: "Give the file name",
		},
	}

	app.Action = func(c *cli.Context) error {
		err := termui.Init()
		if err != nil {
			return err
		}
		defer termui.Close()

		heading := termui.NewPar("GO and Learn!")
		heading.Height = 2
		heading.Width = 20
		heading.Border = false

		g := termui.NewGauge()
		g.Percent = 50
		g.Width = 100
		g.Height = 3
		g.Label = ""
		g.Y = 3
		g.BorderLabel = "GOaL Status 50%"

		courses := []string{
			"[0]Hello World!",
			"[1]2 + 2 = 4",
		}
		ls := termui.NewList()
		ls.Items = courses
		ls.ItemFgColor = termui.ColorYellow
		ls.BorderLabel = "GOaL Courses"
		ls.Height = 12
		ls.Width = 25
		ls.Y = 6
		ls.X = 4

		termui.Render(heading, g, ls)
		termui.Handle("/sys/kbd/q", func(termui.Event) {
			termui.StopLoop()
		})
		termui.Loop()
		return nil
	}

	app.Run(os.Args)
}
