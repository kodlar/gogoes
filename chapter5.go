package main

import (
	"fmt"
)

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex5()
	//ex6()
	ex7()
}

func ex1() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])
}

func ex2() {
	var x [3]float64
	x[0] = 94
	x[1] = 60
	x[2] = 45

	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
	//i klasik başlangıç noktası yani current position i içerde kullanılmadığı için yerine _ koyuyoruz. _ anlamı bu değişkeni kullanmana gerek yok mesajı veriyor
	//value x[i] oluyor
	//range loop edeceğimiz nesneyi temsil ediyor.
	var t float64
	for _, value := range x {
		t += value
	}
	fmt.Println(t / float64(len(x)))

}

//map unordered list
func ex3() {

	//arrays append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	//arrays copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	//maps
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 150
	fmt.Println(y[1])

	fmt.Println(len(y))

	//silmek istersek
	delete(y, 1)

	fmt.Println(len(y))

	elements := make(map[string]string)
	elements["H"] = "hidrojen"
	elements["He"] = "helyum"
	fmt.Println(elements["H"])
	fmt.Println(elements["o"])

	//bu element var mı yok mu ona bakıyoruz.
	name, ok := elements["o"]
	fmt.Println(name, ok)
	//bunun sonucu false gelir
	//element var mı yok mu
	if name, ok := elements["o"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("hata böyle bir element yok!!!")
	}

	if name, ok := elements["H"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("hata böyle bir element yok!!!")
	}
	//map yaratmanın alternatifi object initilizer
	/*
		elem := map[string]string{
			"h":  "hidrojen",
			"he": "helyum",
		}
	*/

	iller := map[string]map[string]string{
		"istanbul": map[string]string{
			"plaka": "34",
			"nüfus": "4 milyon",
		},
		"ankara": map[string]string{
			"plaka": "06",
			"nüfus": "2 milyon",
		},
	}
	if el, ok := iller["ankara"]; ok {
		fmt.Println(el["plaka"], el["nüfus"])
	}
}

func ex4() {
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1[3])
}

func ex5() {
	slice2 := make([]int, 3, 9)
	fmt.Println(len(slice2))
}

func ex6() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	a := x[2:5]
	fmt.Println(a)
}

func ex7() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	min := x[0]

	for _, value := range x {
		if value < min {
			min = value
		}
	}
	//kopya https://socketloop.com/tutorials/golang-find-smallest-number-in-array
	fmt.Println("Bu listedeki en küçük rakam: ", min)
}
