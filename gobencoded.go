package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Employee struct {
	FirstName string `gob:"firstName"`
	 LastName string `gob:"lastName"`
}

type Employees struct {
E []Employee `gob:"employees"`
}

func main() {
	// John Doe, Anna Smith, Peter Jones
	var e Employees
	for i := 1; i < 1000; i++ {
		e.E = append(e.E, Employee{ "John", fmt.Sprintf("Doe the %dth",i)})
	}
	err := gob.NewEncoder(os.Stdout).Encode(&e)
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
}

