package main

import "fmt"

func isMatch(Map [3][3]byte) int {
	for i := 0; i < 3; i += 1 {
		if Map[i][0] == Map[i][1] && Map[i][0] == Map[i][2] {
			if Map[i][0] == 'X' {
				return -1
			} else if Map[i][0] == 'O' {
				return 1
			}
		}
		if Map[0][i] == Map[1][i] && Map[0][i] == Map[2][i] {
			if Map[0][i] == 'X' {
				return -1
			} else if Map[0][i] == 'O' {
				return 1
			}
		}
	}
	if Map[0][0] == Map[1][1] && Map[0][0] == Map[2][2] {
		if Map[0][0] == 'X' {
			return -1
		} else if Map[0][0] == 'O' {
			return 1
		}
	}
	if Map[0][2] == Map[1][1] && Map[0][2] == Map[2][0] {
		if Map[0][2] == 'X' {
			return -1
		} else if Map[0][2] == 'O' {
			return 1
		}
	}
	return 0
}

func main() {
	var Map [3][3]byte
	var tmp string
	for i := 0; i < 3; i += 1 {
		fmt.Scan(&tmp)
		for j := 0; j < 3; j += 1 {
			Map[i][j] = byte(tmp[j])
		}
	}
	winner := isMatch(Map)
	switch winner {
	case -1:
		fmt.Println("X wins")
	case 1:
		fmt.Println("O wins")
	default:
		fmt.Println("Its a tie")
	}
}
