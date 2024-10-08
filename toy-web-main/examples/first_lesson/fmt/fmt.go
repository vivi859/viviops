package main

import (
	"fmt"
)

func main() {
	name := "Tom"
	age := 17
	// 这个 API 是返回字符串的，所以大多数时候我们都是用这个
	str := fmt.Sprintf("hello, I am %s, I am %d years old 111\n", name, age)
	println(str)
	println("testvi")
	println(name)

	// 这个是直接输出，一般简单程序 DEBUG 会用它输出到一些信息到控制台
	fmt.Printf("hello, I am %s, I am %d years old \n", name, age)

	replaceHolder()

	var src = 2022
	fmt.Printf("src=%v ,src&=%v  \n testvi\n", src, &src)
	VariablesConsta(&src)
	//func
	fmt.Printf("func_testing\n")
	Function()
	Defer()
}

// defer延迟执行，下面是个闭包 闭包中变量i 是可以累加的 会被记住
func DeferUil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本地调用接受到n=%v 次调用\n", n)
		i++
		fmt.Printf("匿名工具函数被第i=%v 次调用 \n", i)
		return i
	}
}

func Defer() int {
	f := DeferUil()
	//f(1)
	//return f(2)

	//f(3)先被第一次调用，然后在执行f2
	//defer f(2)
	//return f(3)

	//f(f(4)) f(4)第一次会被先调用，返回值是1，然后再f5，然后执行返回值f(1)
	defer f(f(4))
	return f(5)

	defer f(4)
	defer f(5)
	return f(6)
	//执行6,5,4
}

func replaceHolder() {
	u := &user{
		Name: "Tom",
		Age:  17,
	}
	fmt.Printf("v => %v \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)
}

type user struct {
	Name string
	Age  int
}

func VariablesConsta(n *int) {
	var _u int //默认指0
	fmt.Printf("_u=%v \n", _u)
	fmt.Printf("*nnnn=%v ,&nnn=%v,nnnn=%v\n", *n, &n, n)
}

// 函数
func Getsum(n1 int, n2 int) (int, int) {
	return (n1 + n2), (n1 - n2)

}

func Function() {
	res1, res2 := Getsum(5, 3)
	fmt.Printf("res1=%v,res2=%v \n", res1, res2)
	//函数也是有地址
	fmt.Printf("getsum=%v ,getsum=%T \n", Getsum, Getsum)
}
