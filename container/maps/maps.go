package main

import "fmt"

func main() {
	fmt.Println("Creating map")
	m1 := map[string]string{
		"name":  "james",
		"study": "English",
		"age":   "25",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name, ok := m1["name"]
	fmt.Println(name, ok)

	if name, ok := m1["nam1"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	age, ok := m1["age"]
	fmt.Println(age, ok)

	delete(m1, "age")
	age, ok = m1["age"]
	fmt.Println(age, ok)
}
