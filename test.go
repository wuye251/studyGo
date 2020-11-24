package main // 声明 main 包，表明当前是一个可执行程序

import "fmt" // 导入内置 fmt 包

func main(){ // main函数，是程序执行的入口
	// fmt.Println("Hello World!") // 在终端打印 Hello World!

	// 1. 基本定义
	// var a string = "Go go go"
	// fmt.Println(a)

	// 2. 多值定义
	// var b, c = 1,2
	// fmt.Println(b, c)

	// 3.简短定义
	// b:=123
	// fmt.Println(a)

	// 4.变量声明
	// var a string

	// 5. 定义的值必须使用
	// var a int
	// a=123 
	// fmt.Println("a", a)   

	// 6. 9x9乘法表
	// row:=9
	// multiply(row)

	// 7. 指针
	// var a int = 123
	// var b *int=&a
	// var c **int=&b
	// fmt.Printf("指向a的指针 %x\n",b)
	// fmt.Printf("指向a的指针的指针 %x\n",c)
	// fmt.Printf("指向a的指针的变量 %d\n",*b)
	// fmt.Printf("指向a的指针的指针的变量 %d\n",**c)

	// 8. 结构体
	// type Book struct {
	//  title string
	//  author string
	//  subject string
	//  book_id int
	// }
	// var tmp Book
	// tmp.title = "书的名字"   
	// tmp.author = "书的作者"  
	// tmp.subject = "书的科目"     
	// tmp.book_id = 1
	// fmt.Println(tmp) 
}

func multiply(row int) {
	var ret int
	for i:=1; i<=row; i++ {
		for j:=1; j<=i; j++ {
			ret = i * j
			fmt.Printf("%d",ret)
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}
