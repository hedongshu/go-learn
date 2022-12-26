package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for index, num := range nums {
		fmt.Println("index:", index, "num:", num)
		sum += num
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s", k, v)
		fmt.Println(' ')
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
