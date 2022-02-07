package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello world\n ")

	// 変数宣言
	num := 1
	fmt.Println(num)

	// 型宣言
	num01 := 123
	var num02 int = 1234
	num03 := 1.23
	var num04 float64 = 1.234

	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))

	// 文字列
	string_a := "hello world"

	fmt.Println(reflect.TypeOf(string_a))

	// bool型
	a := 10
	b := 1
	var num_bool bool = a > b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

	// 配列
	arr := [3]string{"sato", "suzuki", "takahashi"}

	fmt.Println(arr[0])

	// 要素数省略可能
	arr2 := [...]string{"sato", "suzuki", "takahashi"}

	fmt.Println(arr2[1])

	// 2次元配列
	arr3 := [2][2]string{
		{"sato", "suzuki"},
		{"takahashi", "tanaka"},
	}

	fmt.Println(arr3[1][0])

	// 条件分岐
	x := 10
	y := 12

	if age := x + y; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

}
