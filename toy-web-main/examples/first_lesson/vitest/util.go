package vitest

import "fmt"

// defer延迟执行，下面是个闭包 闭包中变量i 是可以累加的 会被记住
i := 0
var F =func(s string) int {
	fmt.Printf("本地被%s=%v 次调用\n", s)
	i++
	fmt.Printf("匿名工具函数被第i=%v 次调用 \n", i)
	return i
}

