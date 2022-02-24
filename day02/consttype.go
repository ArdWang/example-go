package main

import "fmt"

// 常量用来做枚举类型
//const (
//	Unknown = 0
//	Female = 1
//	Male = 2
//)

//const (
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)

/*func main(){
	//const LENGTH = 10
	//const WIDTH = 5
	//
	//var area int
	//const a,b,c = 1,false,"str" // 多重赋值
	//
	//area = LENGTH * WIDTH
	//fmt.Printf("面积为:%d",area)
	//println()
	//println(a,b,c)
	//println(a, b ,c)

	const(
		a = iota // 0
		b    //1
		c //2
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a,b,c,d,e,f,g,h,i)


}*/

const (
	i=1<<iota
	j=3<<iota
	k
	l
)

func main(){
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}