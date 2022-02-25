package main

import "fmt"

// 全局变量
var g int

func main2(){
	var a, b ,c int

	a = 10
	b = 20
	c = a+b
	fmt.Printf("result: a = %d, b = %d and g = %d\n",a,b,c)
}



func mainx(){
	var a, b int

	a = 10
	b = 20
	g = a+b
	fmt.Printf("result: a = %d, b = %d and g = %d\n",a,b,g)
}

// 声明全局变量

var k int = 20

func main3(){
	var g int = 10
	fmt.Printf("result: g = %d\n",g)
}


// 形式参数

var a int = 20

func main(){
	var a = 10
	var b = 20
	var c = 0
	fmt.Printf("main()函数中 a = %d\n",a)
	c = sum(a, b)
	fmt.Printf("main()函数中 c = %d\n",c)
}

func sum(a, b int) int{
	fmt.Printf("sum()函数中 a = %d\n",a)
	fmt.Printf("sum()函数中 b = %d\n",b)
	return a+b
}


