package main

import (
	"bufio"
	"io"
	"strings"

	"gonum.org/v1/gonum/floats"
)

func parseInput(r io.Reader, missing string) (vals []float64, weights []float64, err error) {
	sc := bufio.NewScanner(r)

	if missing == "" {
		missing = "missing"
	}

	for sc.Scan() {
		input := strings.TrimSpace(sc.Text())
		if input == "" {
			continue
		}

		val, weight, scanErr := floats.ParseWithNA(input, missing)
		err = scanErr
		if err != nil {
			return
		}

		vals = append(vals, val)
		weights = append(weights, weight)
	}

	err = sc.Err()
	if err != nil {
		return
	}

	return
}
