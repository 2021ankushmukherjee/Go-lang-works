// This is a program Which read a excel file and collect data from phone no cell  and also collect data from json  phone no, then compare that field data is same or not

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xuri/excelize/v2" // import package for read excel file
)

func main() {

	// Read Excel file
	f, err := excelize.OpenFile("data.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	namec1, err := f.GetCellValue("Sheet1", "A2")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(namec1)

	agec2, err := f.GetCellValue("Sheet1", "B2")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(agec2)

	phonec3, err := f.GetCellValue("Sheet1", "C2")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(phonec3)

	exceldata := User{namec1, agec2, phonec3}

	fmt.Println("From excel file:", exceldata)

	// Read json file

	var jsonuser User

	jsondata, _ := ioutil.ReadFile("data.json")

	//fmt.Println(jsondata)

	dson := json.Unmarshal([]byte(jsondata), &jsonuser)

	if dson != nil {

		fmt.Println(dson)
	}
	fmt.Println("From json file: ", jsonuser)

	// checking
	if exceldata.Phone == jsonuser.Phone {

		fmt.Println("Number is matched")
	} else {

		fmt.Println("Number not matched")

	}

}

type User struct {
	Name  string `json: "name"`
	Age   string `json: "age"`
	Phone string `json: "phone"`
}
