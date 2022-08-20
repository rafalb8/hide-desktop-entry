package main

import "strconv"

func isNum(n string) bool {
	for _, r := range n {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func toNum(n string) *int {
	i, err := strconv.Atoi(n)
	if err != nil {
		return nil
	}
	return &i
}

func listDifference(a, b []string) []string {
	if len(a) < len(b) {
		a, b = b, a
	}

	// now a > b

	// b to map
	set := make(map[string]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}

	// iterate over a
	diff := make([]string, 0, len(a)-len(b))
	for _, v := range a {
		if _, exists := set[v]; !exists {
			diff = append(diff, v)
		}
	}

	return diff
}
