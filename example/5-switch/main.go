package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend", time.Now().Weekday())
	default:
		fmt.Println("It's a weekday", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon", t.Hour())
	default:
		fmt.Println("It's after noon", t.Hour())
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool", t)
		case int:
			fmt.Println("I'm an int", t)
		default:
			fmt.Println("Don't know type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
