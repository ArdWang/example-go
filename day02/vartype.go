package main

import "fmt"

var x, y int
var( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"


func main()  {

	//var a string = "Runoob"
	//fmt.Println(a)
	//
	//var b, c int = 1, 2
	//fmt.Println(b, c)

	// 声明一个变量并初始化
	// var a = "RUNOOB"
	// fmt.Println(a)

	// 没有初始化就为零值
	// var b int
	// fmt.Println(b)

	// bool 零值为 false
	// var c bool
	// fmt.Println(c)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n",i, f, b, s)

	// 第二种，根据值自行判定变量类型
	// var d = true
	// fmt.Println(d)

	// 第三种，:= 声明变量
	// f := "Runoob"
	// fmt.Println(f)

	// g, h := 123, "hello"
	// println(x, y, a, b, c, d, e, f, g, h)

	_,numb,strs := numbers() // 只获取函数返回值的后两个
	fmt.Println(numb, strs)


}


// 一个可以返回多个值的函数

func numbers()(int, int, string){
	a, b, c := 1, 2, "str"
	return a,b,c
}