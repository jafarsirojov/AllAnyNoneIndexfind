package main

import (
	"fmt"
)

var db = []account{
	{id: 1, name: "Petya", balance: 100_000},
	{id: 2, name: "Vasya", balance: 90_000},
	{id: 3, name: "Vanya", balance: 450_000},
	{id: 4, name: "Masha", balance: 80_000},
	{id: 5, name: "Nastya", balance: 180_000},
	{id: 6, name: "Sasha", balance: 50_000},
	{id: 7, name: "Oleg", balance: 1_500_000},
}

func ExampleFilterBy_NoResult() {
	result := filterBy(db, func(item account) bool {
		return item.balance < 10_000
	})
	fmt.Println(result)
	// Output: []
}

func ExampleFilterBy_OneResult() {
	result := filterBy(db, func(item account) bool {
		return item.balance > 1_000_000
	})
	fmt.Println(result)
	// Output: [{7 Oleg 1500000}]
}

func ExampleFilterBy_AllResult() {
	result := filterBy(db, func(item account) bool {
		return item.balance > 40_000
	})
	fmt.Println(result)
	// Output: [{1 Petya 100000} {2 Vasya 90000} {3 Vanya 450000} {4 Masha 80000} {5 Nastya 180000} {6 Sasha 50000} {7 Oleg 1500000}]
}

func ExampleIndex_NoResult() {
	result := Index(db, func(item account) bool {
		return item.balance < 10_000
	})
	fmt.Println(result)
	// Output: -1
}

func ExampleIndex_HasResult() {
	result := Index(db, func(item account) bool {
		return item.balance > 100_000
	})
	fmt.Println(result)
	// Output: 2
}

func ExampleAll_False() {
	result := All(db, func(item account) bool {
		return item.balance > 2_000_000
	})
	fmt.Println(result)
	// Output: false
}

func ExampleAll_True() {
	result := All(db, func(item account) bool {
		return item.balance > 20_000
	})
	fmt.Println(result)
	// Output: true
}

func ExampleAny_False() {
	result := Any(db, func(item account) bool {
		return item.balance < 10_000
	})
	fmt.Println(result)
	// Output: false
}

func ExampleAny_True() {
	result := Any(db, func(item account) bool {
		return item.balance > 500_000
	})
	fmt.Println(result)
	// Output: true
}

func ExampleNone_False() {
	result := None(db, func(item account) bool {
		return item.balance < 10_000
	})
	fmt.Println(result)
	// Output: true
}

func ExampleNone_True() {
	result := None(db, func(item account) bool {
		return item.balance > 500_000
	})
	fmt.Println(result)
	// Output: false
}

func ExampleFind_NoResult() {
	result, ok := Find(db, func(item account) bool {
		return item.balance < 10_000
	})

	fmt.Println(result, ok)
	// Output: {0  0} Not found
}

func ExampleFind_HasResult() {
	result, ok := Find(db, func(item account) bool {
		return item.balance > 70_000
	})

	fmt.Println(result, ok)
	// Output: {1 Petya 100000} <nil>
}

func ExampleFind2_NoResult() {
	result, ok := Find2(db, func(item account) bool {
		return item.balance < 10_000
	})

	fmt.Println(result, ok)
	// Output: {0  0} Not found
}

func ExampleFind2_HasResult() {
	result, ok := Find2(db, func(item account) bool {
		return item.balance > 70_000
	})

	fmt.Println(result, ok)
	// Output: {1 Petya 100000} <nil>
}

func ExampleFind3_HasResult() {
	result := Find3NoError(db, func(item account) bool {
		return item.balance > 70_000
	})

	fmt.Println(result)
	// Output: {1 Petya 100000}
}
