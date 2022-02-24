package main

import (
	"fmt"
	"strings"
)


// 定义常量
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构类型声明
type gopher struct {}

// 接口的声明
type golang interface {}


func main(){
	//fmt.Println("Hello,World!")
	//fmt.Println(myMath.Add(1,1))
	////+加号连接
	//fmt.Println("Google"+"Runoob")
	//
	//// %d 表示整型数字，%s表示字符串
	//var stockcode = 123
	//var enddate = "2020-12-31"
	//var url = "Code=%d&endDate=%s"
	//var target_url = fmt.Sprintf(url,stockcode,enddate)
	//fmt.Println(target_url)

	// fmt.Print("a","b",1,2,3,"c","d","\n")
	// fmt.Println("a","b",1,2,3,"c","d")
	// fmt.Printf("ab %d %d %d cd\n",1,2,3)

	//if err := percent(30,70,90,160); err != nil{
	//	fmt.Println(err)
	//}

	str := "这里是 www\n.runoob\n.com"
	fmt.Println("-------原字符串--------")
	fmt.Println(str)

	// 去除空格
	str = strings.Replace(str," ","",-1)
	// 去除换行符
	str = strings.Replace(str,"\n","",-1)
	fmt.Println("----------去除空格与换行后----------")
	fmt.Println(str)

}

func percent(i ...int) error {
	for _, n := range i{
		if n > 100 {
			return fmt.Errorf("数值 %d 超出范围(100)",n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}
