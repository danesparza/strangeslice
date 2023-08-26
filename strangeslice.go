package strangeslice

import (
	"sort"
	"strconv"
	"strings"
)

// StrRangeToInts parses a string of number ranges to a slice of ints
func StrRangeToInts(source string) []int {
	retval := []int{}

	//	First, split based on commas
	commaSplits := strings.Split(source, ",")

	//	Create a temporary map to weed out duplicates
	allItems := map[int]bool{}

	//	Evaluate each item:
	for _, item := range commaSplits {

		// 	- Trim space.  Is there anything left?  If not, continue
		item := strings.TrimSpace(item)
		if len(item) < 1 {
			continue
		}

		// 	- If it has a dash, it's a range.
		if strings.Contains(item, "-") {
			//	Split based on dash
			dashSplits := strings.Split(item, "-")

			//	Make sure we have 2 items
			if len(dashSplits) < 2 {
				continue // Just go to the next item
			}

			//	Trim space for each item
			leftItem := strings.TrimSpace(dashSplits[0])
			rightItem := strings.TrimSpace(dashSplits[1])

			//	Try to parse start/end numbers.
			//	If I can't, just continue
			numLeft, err := strconv.Atoi(leftItem)
			if err != nil {
				continue // Just go to the next item
			}

			numRight, err := strconv.Atoi(rightItem)
			if err != nil {
				continue // Just go to the next item
			}

			//	If they are the same number, add it and continue
			if numLeft == numRight {
				allItems[numLeft] = true
				continue // Go to the next item (and stop wasting my time!)
			}

			//	See what direction we need to go in:
			if numLeft < numRight {
				for i := numLeft; i <= numRight; i++ {
					allItems[i] = true
				}
			} else {
				for i := numRight; i <= numLeft; i++ {
					allItems[i] = true
				}
			}

		} else {
			//	If it doesn't have a dash, try to parse as a number
			numToAdd, err := strconv.Atoi(item)
			if err != nil {
				continue // Just go to the next item
			}

			//	--- Add it to the map
			allItems[numToAdd] = true
		}
	}

	//	Cycle through the map and add the results to the retval
	for key, _ := range allItems {
		retval = append(retval, key)
	}

	//	Sort the slice
	sort.Slice(retval, func(i, j int) bool {
		return retval[i] < retval[j]
	})

	return retval
}
