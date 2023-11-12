package bluelion

import (
	"sort"
	"strings"
)

type SortOptions int

const (
	CaseSensitive   SortOptions = iota // Default behavior is case-sensitive
	CaseInsensitive                    // Use this option for case-insensitive sorting
)

func sortSlice(lines *[]string, options ...SortOptions) {
	caseSensitive := true // Default to case-sensitive

	if len(options) > 0 {
		option := options[0]

		switch option {
		case CaseInsensitive:
			caseSensitive = false
		case CaseSensitive:
			// No need to change anything, already case-sensitive
		}
	}

	sort.Slice(*lines, func(i, j int) bool {
		if caseSensitive {
			return (*lines)[i] < (*lines)[j]
		}
		return strings.ToLower((*lines)[i]) < strings.ToLower((*lines)[j])
	})
}
