package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readHeader(in *bufio.Scanner) (int, int) {
	in.Scan()
	headerLine := in.Text()
	splitted := strings.Split(string(headerLine), " ")
	w, err_w := strconv.Atoi(splitted[0])
	h, err_h := strconv.Atoi(splitted[1])
	if err_w != nil || err_h != nil {
		panic("cannot convert to number!!\n")
	}
	return w, h
}

func readMatrix(in *bufio.Scanner, bitMap *[][]bool) {
	for i := 0; i != len(*bitMap); i += 1 {
		in.Scan()
		line := in.Text()
		for j := 0; j != len((*bitMap)[0]); j += 1 {
			(*bitMap)[i][j] = false
			if line[j] == '1' {
				(*bitMap)[i][j] = true
			}
		}
	}
}

func colorIt(bitMap *[][]bool, w, h, i, j int) {
	if i < 0 || i >= h || j < 0 || j >= w || (!(*bitMap)[i][j]) {
		return
	}
	(*bitMap)[i][j] = false
	for x := -1; x < 2; x += 1 {
		for y := -1; y < 2; y += 1 {
			colorIt(bitMap, w, h, i+x, j+y)
		}
	}

}

func count(bitMap *[][]bool, w, h int) int {
	counter := 0
	for i := 0; i < h; i += 1 {
		for j := 0; j < w; j += 1 {
			if (*bitMap)[i][j] == true {
				counter += 1
				colorIt(bitMap, w, h, i, j)
				//				draw(bitMap)
			}
		}
	}

	return counter
}

func draw(bitMap *[][]bool) {
	for _, x := range *bitMap {
		for _, y := range x {
			if y {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n\n\n")
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	w, h := readHeader(in)
	if w == 0 || h == 0 {
		fmt.Println("0")
	}
	bitMap := make([][]bool, h)
	for i, _ := range bitMap {
		bitMap[i] = make([]bool, w)
	}
	readMatrix(in, &bitMap)
	num := count(&bitMap, w, h)
	fmt.Println(num)
}
