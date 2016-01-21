package main

import "fmt"

func gcd(a, b int) int {
	tmp := 1
	if a < b {
		tmp = a
		a = b
		b = tmp
	}
	for tmp != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}

func gcdList(nums []int) int {
	gcdNum := nums[0]
	for i := 1; i < len(nums); i += 1 {
		gcdNum = gcd(gcdNum, nums[i])
	}
	return gcdNum
}

func main() {
	var db int
	x, err := fmt.Scan(&db)
	nums := make([]int, db)
	var num int
	for i := 0; i < db; i += 1 {
		fmt.Scan(&num)
		nums[i] = num
	}
	fmt.Println(gcdList(nums))
}
