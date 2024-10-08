package main

import (
	"fmt"
)

func main() {

	//func
	fmt.Printf("func_testing\n")
	Defer()

	aa := F("aaaa")
	fmt.Printf("aaaa %v", aa)

	Slice()

	key := Selectkey("注册", "登录", "退出")
	fmt.Println("接收到的key=", key)

}

// 接受多个参数...
func Selectkey(text ...string) (key int) {
	for i, v := range text {
		fmt.Printf("%d:%s\n", i, v)
	}
	fmt.Printf("请选择数字：\n")
	fmt.Scanln(&key)
	return
}

// 切片
func Slice() {
	arr := [5]int{1, 2, 3, 4, 5}
	var s1 []int = arr[0:5] //前开后闭区间
	fmt.Printf("arr= %v ,s1 =%v", arr, s1)
	s1[1] = 1111
	fmt.Printf("arr= %v ,s1 =%v", arr, s1)
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

func F(s string) int {
	i := 0
	fmt.Printf("init 初始化函数%s 次调用\n", s)
	i++
	fmt.Printf("init 初始化函数=%v 次调用 \n", i)
	return i
}
