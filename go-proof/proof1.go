package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dog struct {
	DogName string `json:"dog_name"`
}

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	HasDog    bool
	HasCat    bool  `json:"-"`
	Dogs      []Dog `json:"dog"`
}

func main() {
	jsonBytes := []byte(`
		{
			"id": 1,
			"first_name": "Leandro"		
		}
	`)

	var me Person

	err := json.Unmarshal(jsonBytes, &me)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Hello", me.FirstName, ". Your ID is:", me.ID)

	var me_too Person

	me_too.ID = 2
	me_too.FirstName = "Gabriel"
	me_too.LastName = "Gallac"
	me_too.HasDog = true
	me_too.Dogs = []Dog{{DogName: "Pia"}, {DogName: "Tesa"}}

	out, _ := json.MarshalIndent(me_too, "", "\t")

	fmt.Println(string(out))
}
