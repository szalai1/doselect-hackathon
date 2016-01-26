package main

import "fmt"

func readData(m *map[int]map[int]int) {
	var db int
	fmt.Scan(&db)
	var from, to, dist int
	for i := 0; i < db; i += 1 {
		fmt.Scanf("%d %d %d", &from, &to, &dist)
		if _, ok := (*m)[from]; !ok {
			(*m)[from] = make(map[int]int)
		}
		(*m)[from][to] = dist
		if _, ok := (*m)[to]; !ok {
			(*m)[to] = make(map[int]int)
		}
		(*m)[to][from] = dist
	}
}

/*
 * In m all int must be >= 0
 */
func dijkstra(m *map[int]map[int]int, source int) map[int]int {
	retVal := make(map[int]int)
	unseen := make(map[int]bool)
	for k, _ := range *m {
		retVal[k] = -1
		unseen[k] = true
	}
	delete(unseen, source)
	retVal[source] = 0
	for len(unseen) != 0 {
		minv, mink := -1, -1
		for k, v := range retVal {
			if minv == -1 || (v != -1 && v > minv) {
				mink, minv = k, v
			}
		}
		delete(unseen, mink)
		for k, v := range (*m)[mink] {
			if retVal[k] != -1 {
				if retVal[k] > v+minv {
					retVal[k] = v + minv
				}
			} else {
				retVal[k] = v + minv
			}
		}
	}
	return retVal
}

func main() {
	routs := make(map[int]map[int]int)
	readData(&routs)
	fmt.Println("routs: ", routs)
	var f1, f2 int
	fmt.Scanf("%d %d", &f1, &f2)
	fmt.Println("f1 f2: ", f1, f2)
	d1 := dijkstra(&routs, f1)
	d2 := dijkstra(&routs, f2)
	fmt.Println(routs, "\n", d1, d2)
}
