package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	appName      = "sts"
	appVersion   = "1.0"
	appUsage     = "pronounced 'stats'\n\n	 Generate simple stats for a stream of numbers"
	appUsageText = `seq 1 10 | sts
	 sts numbers.txt
	 seq 1 10 | sts summary
	 seq 1 10 | sts s
`
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.UsageText = appUsageText
	app.Version = appVersion
	app.HideHelp = true
	app.HideVersion = true

	app.Commands = []cli.Command{
		{
			Name:    "summary",
			Aliases: []string{"s"},
			Usage:   "(default) prints summary statistics for the dataset",
			Action:  Summary,
		},
	}

	app.Action = Summary

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show this help message",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "there was an error computing stats for the input: %v\n", err)
		os.Exit(1)
	}
}
