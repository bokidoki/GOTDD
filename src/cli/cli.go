package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	app = cli.NewApp()
)

func init() {
	app.Action = func(ctx *cli.Context) error {
		fmt.Println("exec action")
		if ctx.GlobalIsSet("bool") {
			fmt.Print("set bool is true")
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "bool",
		},
		cli.StringFlag{
			Name: "string",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "command1",
			Action: func(ctx *cli.Context) {
				fmt.Println("exec command1 action")
			},
		},
		{
			Name: "command2",
			Action: func(ctx *cli.Context) {
				fmt.Println("exec command2 action")
			},
		},
		anotherCmd,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("err is %v", err)
	}
}
