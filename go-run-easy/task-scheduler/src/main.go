package main

import (
	"fmt"
	"sort"
)

/*
brute forse
*/
func merge(listA, listB []interface{}, compare func(interface{}, interface{}) bool) (retVal []interface{}) {
	retVal = make([]interface{}, len(listA)+len(listB))
	indA := 0
	indB := 0
	for i := 0; i < len(retVal); i += 1 {
		if compare(listA[indA], listB[indB]) {
			retVal[i] = listB[indB]
			indB += 1
		} else {
			retVal[i] = listA[indA]
			indA += 1
		}
	}
	return retVal
}

func sortOwn(list []interface{}, compare func(interface{}, interface{}) bool) []interface{} {
	if len(list) <= 1 {
		return list
	} else if len(list) == 2 {
		if compare(list[0], list[1]) {
			return []interface{}{list[1], list[0]}
		}
	}
	partA := sortOwn(list[0:len(list)/2], compare)
	partB := sortOwn(list[len(list)/2:], compare)
	return merge(partA, partB, compare)
}

func compId(a, b Event) bool {
	return a.Id > b.Id
}

func compPrior(a, b Event) bool {
	return a.Prior > b.Prior
}

/*
sophisticated
*/
type Event struct {
	Id    int
	Prior int
}

type ById []Event
type ByPrior []Event

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { tmp := a[i]; a[i] = a[j]; a[j] = tmp }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

func (a ByPrior) Len() int           { return len(a) }
func (a ByPrior) Swap(i, j int)      { tmp := a[i]; a[i] = a[j]; a[j] = tmp }
func (a ByPrior) Less(i, j int) bool { return a[i].Prior > a[j].Prior }

func main() {
	var db int
	fmt.Scan(&db)
	list := make([]Event, db)
	for i := 0; i < db; i += 1 {
		fmt.Scanf("%d %d", &(list[i].Id), &(list[i].Prior))
	}
	sort.Sort(ById(list))
	sort.Stable(ByPrior(list))
	for _, v := range list {
		fmt.Printf("%d", v.Id)
	}

}
