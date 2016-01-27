package main

import "fmt"

func readData() []int {
	var db int
	fmt.Scan(&db)
	retVal := make([]int, db)
	for i := 0; i < db; i += 1 {
		fmt.Scan(&retVal[i])
	}
	return retVal
}

func cumsum(nums []int) []int {
	retVal := make([]int, len(nums))
	tmp := 0
	for i, v := range nums {
		retVal[i] = tmp + v
		tmp += v
	}
	return retVal
}

func maxPos(nums []int) (int, int) {
	v, p := 0, 0
	for i, val := range nums {
		if i == 0 {
			v, p = i, val
			continue
		}
		if v < val {
			v, p = i, val
		}
	}
	return v, p
}

func reverse(nums []int) []int {
	var retVal []int
	retVal = make([]int, len(nums))
	copy(retVal, nums)
	for i, v := range nums {
		retVal[len(retVal)-i-1] = v
	}
	return retVal
}

func maxStreach(nums []int) int {
	cumSum := cumsum(nums)
	rnums := reverse(nums)
	revCumSum := reverse(cumsum(rnums))
	v1, p1 := maxPos(cumSum)
	v2, p2 := maxPos(revCumSum)
	fmt.Println(cumSum)
	fmt.Println(revCumSum)
	fmt.Println(v1, p1)
	fmt.Println(v2, p2)
	if p1 >= p2 {
		return v1 + v2 - cumSum[len(cumSum)-1]
	} else {
		if v1 > v2 {
			return v1
		}
		return v2
	}
}

func main() {
	nums := readData()
	fmt.Println(maxStreach(nums))
}
