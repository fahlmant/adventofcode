package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z int
	id      int
}

type JunctionBoxPair struct {
	box1, box2 int
}

func distance(a, b JunctionBox) float64 {
	xDiff := a.x - b.x
	yDiff := a.y - b.y
	zDiff := a.z - b.z
	return float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)
}

type DSU struct {
	parent map[int]int
	size   map[int]int
	count  int
}

func NewDSU(numBoxes int) *DSU {
	dsu := &DSU{
		parent: make(map[int]int, numBoxes),
		size:   make(map[int]int, numBoxes),
		count:  numBoxes,
	}
	for i := 0; i < numBoxes; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	return d.Find(d.parent[i])
}

func (d *DSU) Union(i int, j int) bool {
	root1 := d.Find(i)
	root2 := d.Find(j)

	if root1 != root2 {

		if d.size[root1] < d.size[root2] {
			root1, root2 = root2, root1
		}
		d.parent[root2] = root1
		d.size[root1] += d.size[root2]
		d.count -= 1
		return true
	}

	return false
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boxes := []JunctionBox{}
	boxId := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		z, _ := strconv.Atoi(splitLine[2])
		box := JunctionBox{x, y, z, boxId}
		boxId += 1
		boxes = append(boxes, box)
	}

	distances := make(map[JunctionBoxPair]float64, len(boxes))
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			distances[JunctionBoxPair{i, j}] = distance(boxes[i], boxes[j])
		}
	}

	keys := slices.Collect(maps.Keys(distances))

	slices.SortFunc(keys, func(a, b JunctionBoxPair) int {
		return cmp.Compare(distances[a], distances[b])
	})

	dsu := NewDSU(len(boxes))

	for i, key := range keys {
		if i > 999 {
			break
		}
		dsu.Union(key.box1, key.box2)
	}

	finalSizes := make(map[int]int)

	for i := 0; i < len(boxes); i++ {
		root := dsu.Find(i)

		finalSizes[root] = dsu.size[root]
	}

	sizeValues := slices.Collect(maps.Values(finalSizes))
	slices.SortFunc(sizeValues, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	total = sizeValues[0] * sizeValues[1] * sizeValues[2]
	fmt.Println(total)
}
