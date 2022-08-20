package main

import (
	"bufio"
	"os"
	"strings"
)

type Range struct {
	From *int
	To   *int
}

func ReadRanges() (ranges []Range) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for _, token := range strings.Split(scanner.Text(), " ") {
		switch {
		case isNum(token):
			ranges = append(ranges, Range{
				From: toNum(token),
			})
		case strings.Contains(token, "-"):
			p := strings.SplitN(token, "-", 2)
			r := Range{
				From: toNum(p[0]),
				To:   toNum(p[1]),
			}
			if r.From == nil || r.To == nil {
				continue
			}
			if *r.From > *r.To {
				r.From, r.To = r.To, r.From
			}
			ranges = append(ranges, r)
		}
	}
	return ranges
}

func RangesSlice[T comparable](ranges []Range, list []T) []T {
	var slice []T
	for _, r := range ranges {
		if r.To == nil {
			slice = append(slice, list[*r.From])
		} else {
			slice = append(slice, list[*r.From:*r.To+1]...)
		}
	}

	return slice
}
