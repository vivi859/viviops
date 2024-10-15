package util

import "fmt"

// defer延迟执行，下面是个闭包 闭包中变量i 是可以累加的 会被记住
var F = func(s string) int {
	fmt.Printf("本地被%s=%v 次调用\n", s)
	i := 0
	i++
	fmt.Printf("匿名工具函数被第i=%v 次调用 \n", i)
	return i
}

func Math(a float32, b float32, c byte) float32 {
	ai := a
	bi := b
	ci := c
	var res float32
	switch ci {
	case '+':
		res = ai + bi
	case '-':
		res = ai - bi
	case '*':
		res = ai * bi
	case '/':
		res = ai / bi
	default:
		fmt.Println("符号不对")
	}
	return res
}
