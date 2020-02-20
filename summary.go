package main

import (
	"io"
	"log"
	"os"
	"sort"
	"text/template"

	"github.com/urfave/cli"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

const outputTemplate = `count: {{.Count}}
min:   {{.Min}}
max:   {{.Max}}
sum:   {{.Sum}}

mean:   {{.Mean}}
stddev: {{.StdDevSample}}

q1:     {{.QuartileOne}}
median: {{.Median}}
q3:     {{.QuartileThree}}
`

// Summary generates summary statistics for a data set
func Summary(c *cli.Context) error {
	if c.Bool("help") {
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

	vals, _, err := parseInput(reader, "")
	if err != nil {
		return err
	}

	s := summaryStatistics{data: vals}
	s.populate()
	s.print(os.Stdout)

	return nil

}

type summaryStatistics struct {
	data          []float64
	Count         int
	Min           float64
	Max           float64
	Mean          float64
	Sum           float64
	QuartileOne   float64
	Median        float64
	QuartileThree float64
	StdDevSample  float64
}

func (s *summaryStatistics) populate() {
	// Quantile requires a sorted data set. Doing so at the outset.
	sort.Float64s(s.data)

	s.Count = len(s.data)
	s.Sum = floats.Sum(s.data)
	s.Min = floats.Min(s.data)
	s.Max = floats.Max(s.data)

	s.Mean = stat.Mean(s.data, nil)
	s.QuartileOne = stat.Quantile(0.25, stat.Empirical, s.data, nil)
	s.Median = stat.Quantile(0.5, stat.Empirical, s.data, nil)
	s.QuartileThree = stat.Quantile(0.75, stat.Empirical, s.data, nil)

	s.StdDevSample = stat.StdDev(s.data, nil)
}

func (s *summaryStatistics) print(w io.Writer) {
	t := template.Must(template.New("output").Parse(outputTemplate))
	err := t.Execute(w, s)
	if err != nil {
		log.Fatal(err)
	}
}
