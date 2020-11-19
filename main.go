package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-tars/tars/new"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("tars new ...")
			return nil
		},
	}

	app.Commands = append(app.Commands, new.Commands()...)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
