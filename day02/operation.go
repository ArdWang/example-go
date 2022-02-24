package main

import "fmt"

// 操作运算符

/*func main(){
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c的值为 %d\n",c)
	c = a - b
	fmt.Printf("第一行 - c的值为 %d\n",c)
	c = a * b
	fmt.Printf("第一行 - c的值为 %d\n",c)
	c = a / b
	fmt.Printf("第一行 - c的值为 %d\n",c)
	a++
	fmt.Printf("第一行 - c的值为 %d\n",a)
	a = 21
	a--
	fmt.Printf("第一行 - c的值为 %d\n",a)
}
*/


// 关系运算符

func main1(){
	var a = 21
	var b = 10

	if(a == b){
		fmt.Printf("第一行 - a 等于 b\n")
	}else{
		fmt.Printf("第一行 - a 不等于 b \n")
	}

	if(a < b){
		fmt.Printf("第二行 - a 小于 b\n")
	}else{
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if(a > b){
		fmt.Printf("第三行 - a 大于 b\n")
	}else{
		fmt.Printf("第一行 - a 不大于 b\n")
	}

	a = 5
	b = 20
	if(a <= b){
		fmt.Printf("第四行 - a 小于等于 b\n")
	}

	if(b >= a){
		fmt.Printf("第四行 - b 大于等于 a\n")
	}

}

// 逻辑运算符
func main2()  {
	var a  = true
	var b = false

	if(a && b){
		fmt.Printf("第一行 - 条件为 true\n")
	}

	if(a || b){
		fmt.Printf("第二行 - 条件为 true\n")
	}

	// 修改a和b的值
	a = false
	b = true
	if(a && b){
		fmt.Printf("第三行 - 条件为 true\n")
	}else{
		fmt.Printf("第三行 - 条件为 false\n")
	}

	if(!(a && b)){
		fmt.Printf("第四行 - 条件为 true\n")
	}

}

// 位运算
func main3(){
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("第一行 - c 的值为 %d\n",c)

	c = a | b
	fmt.Printf("第二行 - c 的值为 %d\n",c)

	c = a ^ b
	fmt.Printf("第三行 - c 的值为 %d\n",c)

	c = a << 2
	fmt.Printf("第四行 - c 的值为 %d\n",c)

	c = a >> 2
	fmt.Printf("第五行 - c 的值为 %d\n",c)
}


// 赋值运算符
func main4(){
	var a = 21
	var c int

	c = a
	fmt.Printf("第1行 - = 运算符实列，c 值为 = %d\n",c)

	c += a
	fmt.Printf("第2行 - += 运算符实列，c 值为 = %d\n",c)

	c -= a
	fmt.Printf("第3行 - -= 运算符实列，c 值为 = %d\n",c)

	c *= a
	fmt.Printf("第4行 - *= 运算符实列，c 值为 = %d\n",c)

	c /= a
	fmt.Printf("第5行 - /= 运算符实列，c 值为 = %d\n",c)

	c = 200

	c <<= 2
	fmt.Printf("第6行 - <<= 运算符实列，c 值为 = %d\n",c)

	c >>= 2
	fmt.Printf("第2行 - >>= 运算符实列，c 值为 = %d\n",c)

	c &= 2
	fmt.Printf("第2行 - &= 运算符实列，c 值为 = %d\n",c)

	c ^= 2
	fmt.Printf("第2行 - ^= 运算符实列，c 值为 = %d\n",c)

	c |= 2
	fmt.Printf("第2行 - |= 运算符实列，c 值为 = %d\n",c)
}

// 其它运算符
func main5(){
	var a = 4
	var b int32
	var c float32
	var ptr *int

	// 运算符实列
	fmt.Printf("第1行 - a 变量类型为 = %T\n",a)
	fmt.Printf("第2行 - b 变量类型为 = %T\n",b)
	fmt.Printf("第3行 - c 变量类型为 = %T\n",c)

	// & 和 * 运算符实列
	ptr = &a
	fmt.Printf("a的值为 %d\n",a)
	fmt.Printf("*ptr 为%d\n",*ptr)
}


// 运算符优先级
func main()  {
	var a = 20
	var b = 10
	var c = 15
	var d = 5
	var e int

	e = (a+b)*c/d
	fmt.Printf("e的值为:%d\n",e)

	e = ((a+b)*c)/d
	fmt.Printf("e的值为:%d\n",e)

	e = a+(b*c)/d
	fmt.Printf("e的值为:%d\n",e)

}
