package main

import "fmt"

func demoFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func flowControl() {
	demoFor()
}
