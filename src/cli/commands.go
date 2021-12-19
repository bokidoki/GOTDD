package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

var (
	anotherCmd = cli.Command{
		Name: "anotherCmd",
		Action: func(ctx *cli.Context) {
			fmt.Printf("is another command from anther file")
		},
	}
)
