package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)
var (
	ToBe  bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)
var c, python, cobol  = true, false, "no!"

func main(){
	println("Hello Go!")

	rand.Seed(time.Now().Unix())

	println(rand.Intn(15))

	println(math.Pi)

	println(addNumbers(3,56))

	println(c, python, cobol)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	sum := 1
	for sum < 1000 {
		if number1 := 7-6; number1 == sum{
			sum += sum
		}
		sum += sum
		println(sum)
	}
	print(personalPow(5,2,20))

}

func personalPow(baseNumber, exponentialNumber, limitNumber float64) float64{
	if v := math.Pow(baseNumber,exponentialNumber); v < limitNumber {
		return v
	}
	return limitNumber
}

func addNumbers(x int, y int) (a,b,c  int){
	a = x+y
	b = x-y
	c = x*y
	return
}
