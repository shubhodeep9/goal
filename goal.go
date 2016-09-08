package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
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
		fmt.Println(c.String("lang"))
		return nil
	}

	app.Run(os.Args)
}
