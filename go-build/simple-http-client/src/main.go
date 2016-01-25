package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var path, w1, w2 string
	fmt.Scan(&path)
	fmt.Scan(&w1)
	fmt.Scan(&w2)
	resp, err := http.Get(path)
	if err != nil {
		panic("cannot get!!\n")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("cannat read body!!\n")
	}
	bodyR := strings.Replace(string(body), w1, w2, -1)
	fmt.Println(bodyR)
}
