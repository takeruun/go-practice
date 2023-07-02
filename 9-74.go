package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	Name      string
// 	Age       int
// 	Nicknames []string
// }

type T struct {
}

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"t,omitempty"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Alias Person
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	p.Name = aux.Name + "!"
	return nil
}

func (p Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return json.Marshal(&struct {
		*Alias
		Name string `json:"name"`
	}{
		Alias: (*Alias)(&p),
		Name:  "Mr." + p.Name,
	})
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Nicknames":["Wednesday","Wed"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
