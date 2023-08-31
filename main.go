package main

import (
	"fmt"
	"math"
)

func main() {
	message := "hello world"
	year := 2023
	fmt.Println(message, year)
	message = "hello world"
	year = 2023
	fmt.Println(message)
	fmt.Println(message)
	fmt.Println(message)
	fmt.Println(message)
	for i := 0; i < 20; i++ {
		fmt.Println(message)

	}
	x := 100
	x = x + 100
	y := 200
	z := sum(x, y)
	fmt.Println(z)

	var a, b, c int
	a = 100
	b = 500
	c = sum(a, b)
	fmt.Println(c)
	fmt.Println(sum(a, b))

	m := luyThua(2, 3)
	fmt.Println(m)
	i := 1
	kiemtrasochan(i)
	slice := []string{"Quan", "Nghi", "Tung"}
	fmt.Println(slice)
	for v, u := range slice {
		fmt.Println(v, u)
	}

	// sum int slice
	slice1 := []int{1, 2, 3, 4, 5}
	u := sumSlice(slice1)
	fmt.Println("sum of slice", u)

	// Slice of string
	slice3 := []string{"Monday", "Tuseday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday1"}
	fmt.Println(slice3)

	for i := 6; i >= 0; i-- {
		fmt.Println(slice3[i])

	}
	reverseSlice(slice3)

	//reverseSlice

	slice4 := []int{21, 12, 34, 43}
	m1 := sumSlice2(slice4)
	fmt.Println("sum of slice:", m1)

	days := []string{"Monday", "Tuseday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(getDays(days))

	slIce := []int{1, 2, 3, 4, 5}
	fmt.Println(slIce)
	ketqua:=checkInclude2(22,slice4)
	fmt.Println(ketqua)
}
func sum(a int, b int) int {
	return a + b
}

/*
1. luy thua
2. kiem tra so chan
3 print "hello world" voi so thu tu
*/
func luyThua(a float64, b float64) float64 {

	return math.Pow(a, b)
}
func kiemtrasochan(t int) {
	if t%2 == 0 {
		fmt.Println(" is even")
	} else {
		fmt.Println(" is odd")
	}
}

func sumSlice(slice []int) int {
	sumSlice := 0
	for _, v := range slice {
		sumSlice = sumSlice + v
		fmt.Println(sumSlice)

	}
	return sumSlice
}

func reverseSlice(slice2 []string) {

	rev_slice2 := []string{}
	for i := range slice2 {
		// reverse the order
		rev_slice2 = append(rev_slice2, slice2[len(slice2)-1-i])
	}
	fmt.Println(rev_slice2)

}
func sumSlice2(slice4 []int) int {
	sumSlice2 := 0
	for _, v := range slice4 {
		sumSlice2 = sumSlice2 + v
		fmt.Println(sumSlice2)
	}
	return sumSlice2
}

// output: Monday:Tuesday:Wednesday

func getDays(slicE []string) string {
	result := ""
	for i, v := range slicE {
		if i == 0 {
			result = result + v
		} else {
			result = result + "_" + v
		}
	}

	fmt.Println(slicE)

	return result
}

func checkInclude1(a int, slice []int) bool {
	result:=false
	for _, v := range slice {
		if a == v {
			result= true
			break
		} 
	}
	return result

}

func checkInclude2(a int, slice []int) bool {
	for _, v := range slice {
		if a == v {
		return true
		} 
	}
	return false

}
