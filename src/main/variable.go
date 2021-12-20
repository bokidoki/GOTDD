package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

// "fmt"
// "os"

var a string

func main() {
	// fileinfo, err := os.Stat(filepath.Join("variables.go"))
	// fmt.Printf("fileinfo is %#v, err is %s", fileinfo, err)

	// args := os.Args
	// fmt.Printf("args are %v", args)

	// args := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("从下标2开始截取args得到的数组为：%v", args[2:])

	// dlgo := flag.Bool("dlgo", false, "Download Go and build with it")

	// fmt.Printf("args are %v", os.Args[1:])

	// flag.CommandLine.Parse(os.Args[1:])

	// fmt.Printf("dlgo is %v", *dlgo)

	// ci := os.Getenv("CI")
	// fmt.Printf("ci is %s", ci)

	// a := []int{1, 2}
	// var b []int
	// c := append(b, a...)

	// fmt.Printf("c is: %v", c)

	app := cli.NewApp()

	app.Name = "test cli"
	app.Author = "Lynd"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "StringFlag",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name: "testCommand",
			Action: func(ctx *cli.Context) error {
				firstArg := ctx.Args().First()
				log.Printf("args are: %v", ctx.Args())
				log.Printf("first arg is: %s", firstArg)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "StringFlag",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

func n() { print(a) }

func m() {
	a = "O"
	print(a)
}
