package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "kk",
		"course":  "golang",
		"site":    "imooc",
		"quality": "not bad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}
