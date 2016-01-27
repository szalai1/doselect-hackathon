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
		if _, ok := (*m)[to]; !ok {
			(*m)[to] = make(map[int]int)
		}
		if v, ok := (*m)[from][to]; ok {
			if v > dist {
				(*m)[from][to] = dist
				(*m)[to][from] = dist
			}
		} else {
			(*m)[from][to] = dist
			(*m)[to][from] = dist
		}
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
		if nb, ok := (*m)[source][k]; ok {
			retVal[k] = nb
		}
	}
	delete(unseen, source)
	retVal[source] = 0
	for len(unseen) != 0 {
		minv, mink := -1, -1
		for k, _ := range unseen {
			if minv == -1 || (retVal[k] != -1 && retVal[k] < minv) {
				mink, minv = k, retVal[k]
			}
		}
		if minv == -1 {
			return retVal
		}
		delete(unseen, mink)
		for k, v := range (*m)[mink] {
			if retVal[k] != -1 && retVal[k] > v+minv {
				retVal[k] = v + minv
			} else {
				retVal[k] = v + minv
			}
		}
	}
	return retVal
}

func findBestFactory(d1, d2 *map[int]int) int {
	sumP := -1
	minP := -1
	for k, v := range *d1 {
		v2 := (*d2)[k]
		if v2 != -1 && v != -1 {
			if sumP == -1 || sumP >= v+v2 {
				var max int
				if v > v2 {
					max = v
				} else {
					max = v2
				}
				if minP == -1 {
					sumP = v + v2
					minP = max
				} else if minP > max {
					minP = max
					sumP = v + v2
				}
			}
		}
	}
	return minP
}

func main() {
	routs := make(map[int]map[int]int)
	readData(&routs)
	var f1, f2 int
	fmt.Scanf("%d %d", &f1, &f2)
	d1 := dijkstra(&routs, f1)
	d2 := dijkstra(&routs, f2)

	fmt.Println(d1, d2, "\n", findBestFactory(&d1, &d2))
}
