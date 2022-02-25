package main

import "fmt"

// 创建返回2个数的最大值
func max(num1, num2 int) int{
	var result int
	if num1 > num2 {
		result = num1
	}else{
		result = num2
	}
	return result
}


func swap(x, y string) (string, string){
	return y,x
}



func main1(){
	//var a = 100
	//var b = 200
	//var ret int
	//ret = max(a, b)
	//fmt.Printf("最大值是:%d\n\n", ret)

	a, b := swap("Google", "Runoob")
	fmt.Print(a, b)

}

func main(){
	println("切片解决九九乘法表")
	var num [] int // 定义切片

	for i:=1;i<10;i++{
		num = append(num,i)
	}

	for i:=1;i<10;i++ {
		for j:=1;j<i+1;j++ {
			value:= num[j-1]*i
			fmt.Printf("%d*%d=%d\t",j,i,value)
		}
		fmt.Println()
	}
}
