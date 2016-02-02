package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var org, repo, date string
	fmt.Scanf("%s %s %s", &org, &repo, &date)
	fmt.Println(org, date, repo)

	resp, _ := http.Get("https://api.github.com/repos/" + org + "/" + repo + "/issues")
	str, _ := ioutil.ReadAll(resp.Body)
	var x []struct {
		Number int
	}
	json.Unmarshal(str, &x)
	fmt.Println(len(x), x[1].Number)
}
