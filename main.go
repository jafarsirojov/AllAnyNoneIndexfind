package main

import "fmt"

type account struct {
	id      int64
	name    string
	balance int64
}

func filterBy(items []account, predicate func(item account) bool) []account {
	result := make([]account, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Index(items []account, predicate func(item account) bool) int {
	tempBool :=true
	predicateIndex := -1
	for index, item := range items {
		if predicate(item) {
			if tempBool{
				predicateIndex=index
				tempBool=false
			}
		}
	}
	return predicateIndex
}
func All(items []account, predicate func(item account) bool) bool {
	return Index(items, func(item account)bool{
		return !predicate(item)
	})==-1
}

func Any(items []account, predicate func(item account) bool) bool {
	return Index(items,predicate) !=-1
}

func None(items []account, predicate func(item account) bool) bool {
	return Index(items, predicate) == -1
}
func Find(items []account, predicate func(item account) bool) (account,error) {
	if Index(items, predicate) != -1{
		return items[(Index(items, predicate))],nil
	}
	return account{},fmt.Errorf("Not found")
}

func Find2(items []account, predicate func(item account) bool) (account,error) {
	tempIndex := Index(items, predicate)
	if tempIndex != -1{
		return items[tempIndex],nil
	}
	return account{},fmt.Errorf("Not found")
}

func Find3NoError(items []account, predicate func(item account) bool) account{
	return items[Index(items,predicate)]
}




func main() {}