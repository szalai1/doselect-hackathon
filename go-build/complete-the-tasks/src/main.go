package main

import (
	"fmt"
	"sort"
)

/* ======================================================================
 * The main idea:
 * 1. tasks <=> interval graph -> perfect graph -> max clique = color number
 *    so our tasks is now to find the color number. but we need independent
 *    set, so a clique in the complementer. but that is not as simple as i
 *    thought
 *
 * 2. Earliest deadline first scheduling is the solution.
 *
 * ======================================================================
 */

type ByStart []Event

func (s ByStart) Len() int           { return len(s) }
func (s ByStart) Less(i, j int) bool { return s[i].Start < s[j].Start }
func (s ByStart) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

/*
 * This function return a "brush", which one is also a function
 *  - if we call the brush with -1 it give us the next unused number (int)
 *  - but if we call it with a number, then that number (color) becames
 *    useable again.
 */

func brushGenerator() func(int) int {
	palette := make([]bool, 1)
	return func(x int) int {
		//fmt.Println(x)
		if x == -1 {
			for i, v := range palette {
				if !v {
					palette[i] = true
					return i
				}
			}
			palette = append(palette, true)
			return len(palette)
		}
		palette[x] = false
		return len(palette)
	}
}

func color(data []Event) int {
	brush := brushGenerator()
	dataEnd := make([]Event, len(data))
	copy(dataEnd, data)
	sort.Sort(ByEnd(dataEnd))
	sort.Sort(ByStart(data))
	fmt.Println("End: ", dataEnd)
	fmt.Println("Star: ", data)
	endI := 0
	time := data[0].Start
	for i, v := range data {
		time = v.Start
		fmt.Println(v, dataEnd[endI])
		for time >= dataEnd[endI].End {
			brush(dataEnd[endI].Color)
			endI += 1
		}
		data[i].Color = brush(-1)
	}
	return brush(0)
}

/* ==================================
 * Earliest deadline first scheduling
 * ==================================
 */

type Event struct {
	Start int
	End   int
	Color int
}

type ByEnd []Event

func (s ByEnd) Len() int           { return len(s) }
func (s ByEnd) Less(i, j int) bool { return s[i].End < s[j].End }
func (s ByEnd) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func readData() []Event {
	var db int
	fmt.Scan(&db)
	retVal := make([]Event, db)
	for i := 0; i < db; i += 1 {
		fmt.Scanf("%d %d", &retVal[i].Start, &retVal[i].End)
	}
	return retVal
}

func EDF(data []Event) int {
	sort.Sort(ByEnd(data))
	counter := 0
	var current Event
	for i, v := range data {
		if i == 0 {
			current = v
			continue
		}
		if v.Start >= current.End {
			current = v
			counter += 1
		}
	}
	return counter + 1
}

func main() {
	data := readData()
	num := EDF(data)
	fmt.Println(num)
}
