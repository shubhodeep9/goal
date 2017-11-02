package main

import (
	"fmt"
	"github.com/gizak/termui"
	"github.com/shubhodeep9/goal/exercises"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strconv"
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
	rotator := 0

	//character code for check mark is \u2714
	highlighted := []string{
		"[[0] Hello World!            [  ]](fg-white,bg-blue)",
		"[[1] 2 + 2 = 4               [  ]](fg-white,bg-blue)",
		"[[2] Server                  [  ]](fg-white,bg-blue)",
		"[[*] EXIT                    [\u2717 ]](fg-white,bg-blue)",
	}

	courses := []string{
		"[0] Hello World!            [  ]",
		"[1] 2 + 2 = 4               [  ]",
		"[2] Server                  [  ]",
		"[*] EXIT                    [\u2717 ]",
	}

	runtimecourses := []string{
		"[[0] Hello World!            [  ]](fg-white,bg-blue)",
		"[1] 2 + 2 = 4               [  ]",
		"[2] Server                  [  ]",
		"[*] EXIT                    [\u2717 ]",
	}
	var FlagSet bool = true

	current, _ := exercises.GetCurrent()

	//main handler for the activity
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			err := termui.Init()
			if err != nil {
				return err
			}
			defer termui.Close()

			heading := termui.NewPar("GO and Learn!")
			heading.Y = 1
			heading.X = 2
			heading.Height = 2
			heading.Width = 20
			heading.Border = false

			g := termui.NewGauge()
			g.Percent = int(current/len(courses)) * 100
			g.Width = 60
			g.Height = 3
			g.Label = ""
			g.X = 18
			g.BorderLabel = "GOaL Status " + strconv.Itoa(g.Percent) + "%"

			ls := termui.NewList()
			ls.Items = runtimecourses
			ls.ItemFgColor = termui.ColorYellow
			ls.BorderFg = termui.ColorBlue
			ls.PaddingLeft = 6
			ls.PaddingTop = 1
			ls.BorderLabel = "GOaL Courses"
			ls.Height = 12
			ls.Width = 50
			ls.Y = 3
			ls.X = 3
			termui.Render(heading, g, ls)
			termui.Handle("/sys/kbd", func(e termui.Event) {
				if e.Path == "/sys/kbd/<enter>" {
					termui.StopLoop()
				}
				if e.Path == "/sys/kbd/q" || e.Path == "/sys/kbd/<escape>" {
					rotator = len(runtimecourses) - 1
					termui.StopLoop()
				}

				//Beginning of iteratable list code
				//Simple incremental algorithm used to
				//highlight list items
				if e.Path == "/sys/kbd/<down>" {
					rotator = rotator + 1
				} else if e.Path == "/sys/kbd/<up>" {
					rotator = rotator - 1
				}
				if rotator < 0 {
					rotator = len(runtimecourses) + rotator
				} else {
					rotator = rotator % len(runtimecourses)
				}
				for i := 0; i < len(runtimecourses); i++ {
					if i == rotator {
						runtimecourses[i] = highlighted[i]
					} else {
						runtimecourses[i] = courses[i]
					}
				}
				ls.Items = runtimecourses
				termui.Render(ls)
			})
			termui.Loop()
			FlagSet = false
		}
		return nil
	}
	app.Run(os.Args)

	// this means that only > goal \
	// was triggered, so we print the choice
	// else the flag is triggered
	if !FlagSet {
		if rotator < len(runtimecourses)-1 {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			exercises.ExerciseGo(rotator)
			fmt.Println(exercises.ExerciseChecker())
		}
	}

}
