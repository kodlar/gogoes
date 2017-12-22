package main

//this is a comment
import "fmt"

//global değişken

const globalVar string = "hasan"

func main() {
	//chapter4for()
	//chapter4foranother()
	chapter4if()
}
func chapter4if() {

	for i := 0; i <= 20; i++ {
		if i%5 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
			switch i {
			case 1:
				fmt.Println(i, "olsa")
			default:
				fmt.Println("hımmm düşünmek istemiyorum")
			}
		}
	}

}
func chapter4foranother() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func chapter4for() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func GlobalSwitch() {
	fmt.Print("Değer giriniz")
	var iput float64
	fmt.Scanf("%f", &iput)
	test := iput
	//fmt.Println(test)

	switch test {
	case 1:
		//fmt.Println("ho")
		FeetToMeters()
		break
	case 2:
		//fmt.Println("hoaha")
		FahrenheitToCelcius()
		break
	}
}

func FeetToMeters() {
	fmt.Print("Enter feet value: ")
	var input float64
	fmt.Scanf("%f", &input)

	formul := input * 0.34048

	fmt.Println(formul, "meter")
}

func FahrenheitToCelcius() {
	fmt.Print("Enter a Fahrenheit value: ")
	var input float64
	fmt.Scanf("%f", &input)

	//var celcius int64
	celcius := (input - 32) * 5 / 9
	fmt.Println(celcius, " celcius")
}

func dinamiktip() {
	o := "test ediyorum bu tipi"
	fmt.Println(o)
}

func biryazdirma() {
	var s string
	s = "hadi bakalım"
	fmt.Println(s)
	s = s + " kolay gelsin"
	fmt.Println(s)
}

func yazdir() {
	x := "hello"
	y := "world"
	fmt.Println(x == y)
}

func hadi() {
	fmt.Println(globalVar)
}

func tras() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func test2() {
	x := 5
	x++
	fmt.Println(x)
}

func getName() {
	//better name convention
	dogname := "çomar"
	fmt.Println("benim köpeğin ismi", dogname)
	fmt.Println("buna bu fonksiyonun dışından eriştim ", globalVar)
}

func chapter2() {

	fmt.Println("Hello world")
	fmt.Println("1 + 1", 1.0+1.0)
	fmt.Println(len("hello world"))
	fmt.Println("hello"[0])
	fmt.Println(32.132 * 42.45)

}
