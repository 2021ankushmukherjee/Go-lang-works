package main

import "fmt"

type Customer struct {
	Name string
	Add  []Address
}

type Address struct {
	Address string
}
type output struct {
	Name string
	Adr  string
}

func main() {

	add1 := Address{

		"E7",
	}
	add2 := Address{

		"EG",
	}

	var cust Customer
	fmt.Print("Input Struct\n")
	cust.Name = "Ankush"
	cust.Add = append(cust.Add, add1)
	cust.Add = append(cust.Add, add2)
	fmt.Println(cust)
	fmt.Print("\n")
	var op output
	fmt.Print("Output Struct\n")
	for i := range cust.Add {
		op.Name = cust.Name
		op.Adr = cust.Add[i].Address
		fmt.Println(op)
	}

}
