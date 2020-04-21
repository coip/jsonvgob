package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	FirstName string `json:"firstName"`
	 LastName string `json:"lastName"`
}

type Employees struct {
E []Employee `json:"employees"`
}

func main() {
	// John Doe, Anna Smith, Peter Jones
	var e Employees
	for i := 1; i < 1000; i++ {
		e.E = append(e.E, Employee{ "John", fmt.Sprintf("Doe the %dth",i)})
	}
	err := json.NewEncoder(os.Stdout).Encode(&e)
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
}

