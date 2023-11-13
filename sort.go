package bluelion

import (
	"sort"
	"strings"
)

type SortOptions int

const (
	CaseSensitive   SortOptions = iota // Default behavior is case-sensitive
	CaseInsensitive                    // Use this option for case-insensitive sorting
	Reverse                            // Use this option to sort in reverse order
)

func sortSlice(lines *[]string, options ...SortOptions) {
	caseSensitive := true // Default to case-sensitive
	reverse := false      // Default to ascending order

	if len(options) > 0 {
		for _, option := range options {
			switch option {
			case CaseInsensitive:
				caseSensitive = false
			case CaseSensitive:
				// No need to change anything, already case-sensitive
			case Reverse:
				reverse = true
			}
		}
	}

	sort.Slice(*lines, func(i, j int) bool {
		if caseSensitive {
			if reverse {
				return (*lines)[i] > (*lines)[j]
			}
			return (*lines)[i] < (*lines)[j]
		}
		if reverse {
			return strings.ToLower((*lines)[i]) > strings.ToLower((*lines)[j])
		}
		return strings.ToLower((*lines)[i]) < strings.ToLower((*lines)[j])
	})
}
