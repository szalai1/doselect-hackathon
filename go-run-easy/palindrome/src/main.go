package main

import "fmt"
import "strings"

func isPalindrome(str string) bool {
	isP := true
	for i := 0; i < len(str)/2; i += 1 {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return isP
}

func main() {
	var db int
	fmt.Scan(&db)
	for db != 0 {
		db -= 1
		var str string
		fmt.Scan(&str)
		str = strings.ToLower(str)
		if isPalindrome(str) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
