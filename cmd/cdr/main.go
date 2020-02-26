package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "CD-R 700MB",
		Usage:   "A static mixtape site generator",
		Version: "v1.0",
		Flags:   []cli.Flag{},
	}

	app.Commands = []*cli.Command{
		&cmdClean,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
