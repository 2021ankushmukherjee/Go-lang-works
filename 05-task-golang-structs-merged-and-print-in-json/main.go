// Merge 3 struct and print in one row in json

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	mcd := companyDetials{"ABC", "1994"}

	cust := CustomerDetails{"Ankush", "Mukherjee"}

	med := EmployeeDetials{"Ram"}

	var m map[string]string

	jmpr, _ := json.Marshal(mcd)
	json.Unmarshal(jmpr, &m)

	jcust, _ := json.Marshal(cust)
	json.Unmarshal(jcust, &m)

	jmtspecd, _ := json.Marshal(med)
	json.Unmarshal(jmtspecd, &m)

	jm, _ := json.Marshal(m)
	fmt.Println(string(jm))

}

type companyDetials struct {
	Name string `json:"Name"`
	Estd string `json:"Estd"`
}

type CustomerDetails struct {
	Name1 string `json:"NAME1"`
	Name2 string `json:"NAME2"`
}

type EmployeeDetials struct {
	Name string `json:"Name"`
}
