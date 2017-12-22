package main

import (
	"fmt"
)

func avarege(xs []float64) float64 {

	total := 0.0
	for _, v := range xs {
		total += v
	}

	//fmt.Println(total / float64(len(xs)))
	return total / float64(len(xs))
	//panic("not implemented")
}

func main2() {

	xs := []float64{98, 93, 77, 82, 73}
	fmt.Println(avarege(xs))

}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func a1() {
	fmt.Println("a1")
}

func a2() {
	fmt.Println("a2")
}

func main() {
	//fmt.Println(f1())
	//main3()
	fmt.Println(add(1, 2, 3))
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))

	he := 0
	arttir := func() int {
		he++
		return he
	}
	fmt.Println(arttir())
	fmt.Println(arttir())
	nexteven := makeEvenGenerator()
	fmt.Println(nexteven())
	fmt.Println(nexteven())
	fmt.Println(nexteven())

	defer a2()
	a1()
	alix()

	x := 5
	pzero(x)
	zero(&x)
	fmt.Println(x)
}

func zero(xPtr *int) {
	*xPtr = 0
}

func pzero(x int) { x = 0 }

func alix() {
	//panic("PANIC")
	str := recover() // this will never happen
	fmt.Println(str)

}

func f1() int {
	return f2()
}

func f2() int {
	return f3()
	//return 1
}

func f3() (r int) {
	r = 1
	return
}

//tuple örneği
func f() (int, int) {
	return 5, 6
}

func main3() {
	x, y := f()
	fmt.Println(x, y)
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
