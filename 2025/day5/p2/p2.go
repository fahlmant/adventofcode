package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file), "\n\n")

	rangeStrings := strings.Split(input[0], "\n")

	ranges := make([]Range, len(rangeStrings))

	for i, v := range rangeStrings {
		rangeSplit := strings.Split(v, "-")
		startOfRange, err := strconv.Atoi(rangeSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		endOfRange, err := strconv.Atoi(rangeSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges[i] = Range{start: startOfRange, end: endOfRange}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	mergedRanges := []Range{}
	for _, ra := range ranges {
		// Base case
		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, ra)
			continue
		}
		// Check the last range in mergedRanges and see if it
		// is separate from the next range
		if mergedRanges[len(mergedRanges)-1].end < ra.start-1 {
			mergedRanges = append(mergedRanges, ra)
		} else {
			if ra.end > mergedRanges[len(mergedRanges)-1].end {
				mergedRanges[len(mergedRanges)-1].end = ra.end
			}

		}
	}

	for _, r := range mergedRanges {
		total += r.end - r.start + 1
	}

	fmt.Println(total)
}
