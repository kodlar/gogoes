package main

import "fmt"

func main() {
	e2()
	e1()
}

func e1() {

	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func e2() {
	z := 0
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
			z++
		}

	}
	fmt.Println(z)

}
