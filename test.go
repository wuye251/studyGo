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
	
	//9.切片
	//s :=[] int {1,2,3}
	//slice1 := s[1:2]
	//slice2 := s[:2]
	//slice3 := s[2:]
	//fmt.Println(s)
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//fmt.Prifmt.Println(slice3)
	//s :=[] int {1,2,3}
	//slice1 := s[1:2]
	//slice1 printSlice(slice1)
	//slice1 slice1 = append(slice1, 99)
	//slice1 printSlice(slice1)
	//slice1 slice1 = append(slice1, 1)
	//slice1 printSlice(slice1)
	//slice1 slice1 = append(slice1,1,1)
	//slice1 printSlice(slice1)
	
	//10.Range
	//nums := []int{0,1,2,3,4}
	//sum := 0
	//for i, num := range nums {
	//    	fmt.Println("cur index:", i)
	//	sum +=num
	//}
	//fmt.Println("sum:", sum)

	//11.map
	//var map1 map[string]string
	//map1 = make(map[string]string)
	//map1["France"] = "巴黎"
	//map1["Italy"]  = "罗马"
	//map1["Japan"]  = "东京"
	//for key, value := range map1 {
	//    fmt.Println(key,"首都是", value)
	//}
	//
	//delete(map1, "France")
	////ok := map1["Japan"]
	////fmt.Println(ok)
	//fmt.Println("----------------")
	//for key := range map1 {
	//    fmt.Println(key,"首都是", map1[key])
	//}

	//12.递归函数
	//num := 7 
	//result := fibonacci(num)
	//fmt.Println(num, "fibonacci num is: ", result)

	//13.错误接口
	result, err := Sqrt(-1)
	if err != nil {
	    fmt.Println(err)
	}
}

type error interface {
  Error() string
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0,error.New("math:square root of negative number")
    } 
}

func fibonacci(num int) int{
  if num < 2 {
    return num
  }
  return fibonacci(num-1) + fibonacci(num-2)
}

func printSlice(x [] int){
 fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
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

