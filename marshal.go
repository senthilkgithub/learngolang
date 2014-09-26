package main

import (
	"fmt"

	"encoding/json"
)

type Envelop map[string]interface{}

type message struct {
	Type, Text string
}

func main() {

	var m = make(Envelop)

	m["message"] = message{"error", "You can't do that!"}

	j, err := json.Marshal(&m)

	if err != nil {

		fmt.Println(err)

	}

	fmt.Printf("Struct: %s\n", m)

	fmt.Printf("JsonIn: %s\n", j)

	var mm map[string]Envelop

	err = json.Unmarshal(j, &mm)

	if err != nil {

		fmt.Println(err)

	}

	var ss Envelop

	ss = mm["message"]

	fmt.Println("MR: ", ss)

}
