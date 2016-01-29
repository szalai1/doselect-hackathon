package main

import "fmt"

func trim(data []int, t int) []int {
	retVal := make([]int, 0, len(data))
	i := 0
	for _, v := range data {
		if v <= t {
			retVal = append(retVal, v)
			i += 1
		}
	}
	return retVal
}

func numberOfCombinations(data []int, t int) int {
	L1 := []int{0}
	for _, v := range data {
		L2 := make([]int, len(L1))
		for i, val := range L1 {
			L2[i] = val + v
		}
		L2 = trim(L2, t)
		L1 = append(L2, L1...)
	}
	counter := 0
	for _, v := range L1 {
		if v == t {
			counter += 1
		}
	}
	return counter
}

func readData() ([]int, int) {
	var db, s int
	fmt.Scanf("%d %d", &db, &s)
	cards := make([]int, db)
	for i := 0; i < db; i += 1 {
		fmt.Scan(&cards[i])
	}
	return cards, s
}

func main() {
	data, t := readData()
	fmt.Println(numberOfCombinations(data, t))

}
