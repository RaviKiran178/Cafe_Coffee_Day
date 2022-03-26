package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var Quantity int
	var Foodtype string
	var DaySummary = []int{}

	run := 0
	Menu_items := `
	----------------------------------------
	             CAFE COFFEE DAY
	----------------------------------------
	S No 	   Items		Price
	----------------------------------------
	1.	C: coffee		40rs/-
	2.	D: dosa			80rs/-
	3.	T: tomato soup		20rs/-
	4.	I: idli			20rs/-
	5.	V: vada			25rs/-
	6.	B: bhature&chole	30rs/-
	7.	P: paneer pakoda	30rs/-
	8.	M: manchurian		90rs/-
	9.	H: hakka noodle		70rs/-
	10.	F: French fries		30rs/-
	11.	J: jalebi		30rs/-
	12.	L: lemonade		15rs/-
	13.	S: spring roll		20rs/-
	14.   END: Exit the day `
	for 1 > 0 {
		fmt.Println(Menu_items, "\n-----------------------------------------------")
		fmt.Print("Enter the Code Of Food: ")
		fmt.Scan(&Foodtype)
		foodtype := strings.ToLower(Foodtype)
		//fmt.Println(foodtype)
		if foodtype == "end" {
			BillSummary(DaySummary)
			break
		} else {
			fmt.Print("Enter the Quantity of Food: ")
			fmt.Scan(&Quantity)
			fmt.Println("---------------------------------------------------")
		}
		daybill := bill(Quantity, foodtype)
		fmt.Println("Bill For:", Foodtype, "(Foodtype) of", Quantity, "(quantity) is Rs.", daybill)
		fmt.Println("---------------------------------------------------")
		DaySummary = append(DaySummary, daybill)
		run += 1
	}
}
func BillSummary(DaySummary []int) {
	Day_income := 0
	dt := time.Now()
	for inc := 0; inc < len(DaySummary); inc++ {
		Day_income = DaySummary[inc] + Day_income
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("Income of (", dt.Format("01-02-2006"), dt.Format("Monday"), ")is Rs.", Day_income, "/-", "\nThan you visit Again...")
}
func bill(Quantity int, foodtype string) int {
	m := map[string]int{
		"c": 40,
		"d": 80,
		"t": 20,
		"i": 20,
		"v": 25,
		"b": 30,
		"p": 30,
		"h": 70,
		"f": 30,
		"j": 30,
		"l": 15,
		"s": 20,
	}
	total := Quantity * m[foodtype]
	return total
}
