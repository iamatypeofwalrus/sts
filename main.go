package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
	"github.com/urfave/cli"
)

const (
	appName      = "sts"
	appVersion   = "1.0"
	appUsage     = "stats\n\nGenerate simple stats for a stream of numbers"
	appUsageText = "sts numbers.txt or cat numbers.txt | sts"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.UsageText = appUsageText
	app.Version = appVersion
	app.HideHelp = true
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show this help message",
		},
	}

	app.Action = do

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "there was an error computing stats for the input: %v\n", err)
		os.Exit(1)
	}
}

func do(c *cli.Context) error {
	if c.Bool("help") || len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	var reader io.Reader
	if c.NArg() > 0 {
		filePath := c.Args().Get(0)
		r, err := os.Open(filePath)
		if err != nil {
			return err
		}

		reader = r
	} else {
		reader = os.Stdin
	}

	input, err := parseInput(reader)
	if err != nil {
		return err
	}

	sts, err := generateStats(input)
	if err != nil {
		return err
	}

	printStats(sts)

	return nil
}

func parseInput(reader io.Reader) ([]float64, error) {
	scanner := bufio.NewScanner(reader)
	input := make([]float64, 0)

	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		if str == "" {
			continue
		}

		float, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}

		input = append(input, float)
	}

	return input, nil
}

func generateStats(data []float64) (*inputStats, error) {
	mean, err := stats.Mean(data)
	if err != nil {
		return nil, err
	}

	min, _ := stats.Min(data)
	max, _ := stats.Max(data)

	stdDev, _ := stats.StandardDeviationSample(data)
	sum, _ := stats.Sum(data)
	quartiles, _ := stats.Quartile(data)

	is := &inputStats{
		count:         len(data),
		mean:          mean,
		quartileOne:   quartiles.Q1,
		median:        quartiles.Q2,
		quartileThree: quartiles.Q3,
		min:           min,
		max:           max,
		stdDev:        stdDev,
		sum:           sum,
	}

	return is, nil
}

func printStats(sts *inputStats) {
	printStat("count", sts.count)
	printStat("min", sts.min)
	printStat("q1", sts.quartileOne)
	printStat("median", sts.median)
	printStat("q3", sts.quartileThree)
	printStat("max", sts.max)
	printStat("mean", sts.mean)
	printStat("sum", sts.sum)
	printStat("stddev", sts.stdDev)
}

func printStat(name string, stat interface{}) {
	fmt.Printf("%s:\t%v\n", name, stat)
}
