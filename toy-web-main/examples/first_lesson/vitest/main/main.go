package main

import (
	"errors"
	//"util"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 这个函数执行完后，这个栈空间就会被删除那么n1这个值就没了，然后下面main打印的就是main本身的n1=10
func test(n int) {
	n1 := n + 1
	fmt.Printf("test n1=%v\n", n1)
}

func digui(n int) {
	if n > 2 {
		n--
		digui(n)
	}
	fmt.Printf("test-digui-n=%v \n", n)
}

func feibona(n int) int {
	if (n == 1) || (n == 2) {
		return 1
	} else {
		return feibona(n-1) + feibona(n-2)
	}
}

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Printf("输入天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

// 闭包里面的n只初始化一次=10，后面每次调用是变化后的值，这个变量会持续变化
func Addr() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// 闭包实践 要求传入jpg，如果后缀是有jpg的，则直接返回，如果灭有，则返回加好后缀的闭包，suffix是函数外部的变量跟返回函数组合成一个闭包
// 闭包传入的参数suffix引用到的这个变量可以一直调用，不需要每次调用，
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name有指定后缀名则返回，没有就要加上
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

// defer的使用 defer后面执行顺序是最近的先执行 n2先执行，在执行N1，先入后出
// defer价值 当函数执行完后，可以释放函数创建的资源,connnect =operdatabase()  defer connect.close()
func defertest(n1 int, n2 int) int {
	defer fmt.Println("defer n1=", n1)
	defer fmt.Println("defer n2=", n2) //把栈放到defer栈，但是值就是当时的值，不会因为后面的++改变
	n1++
	n2++
	res := n1 + n2
	fmt.Println("defer behind res:", res)
	return res
}

// 函数打印金字塔
func jin(ii int) {
	for i := 0; i < ii; i++ {
		for m := 1; m <= (ii - i); m++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// 打印字符
func Print(s string) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Printf("打印字符：%c\n", r[i])
	}

	//字符串转成整形v,整形到字符串是strconv.Itoa
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("转换成字符串%v,类型：%T", n, n)
	}

	//字符串转 []byte：
	var bytes = []byte("hello goo")
	fmt.Printf("bytes=%v\n", bytes)
	//byte转字符串
	str := string([]byte{97, 99})
	fmt.Printf("byte to string str=%v ,T=%T \n", str, str)

	//查找字符串是否存在指定字符串
	b := strings.Contains("seafoof", "foof")
	fmt.Printf("b=%v \n", b)
	num := strings.Count("cheeeess", "e")
	fmt.Printf("num have e:%v \n", num)

}

// time需求
func timeii() {
	i := 0
	for {
		i++
		time.Sleep(time.Millisecond * 1000)
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}

// 得到这个func执行时间
func test001() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

// 错误处理 发生错误panic后，程序会退出，go会抛出一个panic ,在defer中通过recover捕获这个异常。
func errrr() {
	//匿名函数
	defer func() {
		//err := recover()可以分开写也可以写到if里面
		if err := recover(); err != nil {
			fmt.Println("获取到错误打印出来，继续执行程序err=", err)
		}
	}()
	num1 := 100
	num2 := 0
	num3 := num1 / num2
	fmt.Println("num3=", num3)

}

// 自定义错误
func readConf(name string) (err error) {
	if name == "vivitest.conf" {
		return nil
	} else {
		return errors.New("自定义的错误名称")
	}
}
func putConf() {
	err := readConf("vivitest.conf")
	if err != nil {
		panic(err) //如果读取到了这个自定义错误，就输出错误，终止程序
	}
	fmt.Println("没有出现定义错误就继续往下走哦")
}

func main() {

	cityes := make(map[string]string)
	cityes["no1"] = "北京" //如果这个key存在就是更新values，如果不存在就是新增key values
	cityes["no2"] = "北京"
	delete(cityes, "no2")
	fmt.Println("cityes", cityes)
	fmt.Println("长度len",len(cityes))
	//一次性删除所有key，遍历删除，或者 make一个新的	cityes := make(map[string]string)

	students := make(map[string]map[string]string)
	students["no1"] = make(map[string]string, 2)
	students["no1"]["name"] = "小王"
	students["no1"]["age"] = "27"
	fmt.Println("students", students)

	for k1,v1 in :=range students{
		fmt.Println("k1",k1)
		for k2,v2:=range v1{

	}
	}

	////错误处理
	//errrr()
	//fmt.Printf("error错误了不处理不会执行这个哦:")
	////自定义错误
	//putConf()

	// 打印字符
	//Print("hello我是你")
	//now := time.Now()
	//fmt.Printf("now=%v now type=%T ", now, now)
	//fmt.Printf("年%v 月%v 日%v \n", now.Year(), now.Month(), now.Day())
	//time需求 格式化日期 指定这样的日期格式化
	//now := time.Now()
	//fmt.Printf(now.Format("2006/01/02 15:04:05"))
	//timeii()

	////得到这个func执行时间
	//start := time.Now().Unix()
	//test001()
	//end := time.Now().Unix()
	//fmt.Printf("执行test001的时间为秒：%v \n", end-start)

	//金字塔
	//jin(5)

	//defer
	//nn := defertest(1, 2)
	//fmt.Println("nn=", nn)
	//闭包
	//bi := makeSuffix(".jpg")
	//fmt.Println("文件处理后的后缀=", bi("ssss"))
	//fmt.Println("文件处理后的后缀=", bi("aaa.jpg"))

	//nn := feibona(5)
	//fmt.Printf("nn=%v \n", nn)
	//fmt.Println("第8天桃子数量：", peach(9))
	//
	//fmt.Printf("aaaaa")
	//fmt.Printf("bbbb")

	//digui(4)

	////func
	//fmt.Printf("func_testing\n")
	//Defer()
	//
	//aa := F("aaaa")
	//fmt.Printf("aaaa %v\n", aa)
	//
	//n1 := 10
	//test(n1)
	//fmt.Printf("main n1=%v\n", n1)

	//Slice()

	//key := Selectkey("注册", "登录", "退出")
	//fmt.Println("接收到的key=", key)

	//mapp()
	//Sttu()
	//type mesType uint16
	//var u1000 uint16 = 1000 //u1000变量是uint16类型=1000
	//var textmes mesType = mesType(u1000)

	//type myUint16 = uint16
	//var myu16 myUint16 = u1000

	//接口
	//fmt.Printf("\n testing func_interface \n")
	//Interface()
	//Goroutnie()
	//Prints(9)
	//nine()

	//result := util.Math(2.2, 2, '*')
	//fmt.Printf("result=%v \n", result)

}

// 9*9
func nine() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v  ", j, i, j*i)
		}
		fmt.Println()
	}

}

// 打印
func Prints(aa int) {
	for i := 1; i < aa; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 协程 goroutine
var c int
var wg sync.WaitGroup // 用于等待所有协程完成

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%v\n", n)
	c++
}
func Goroutnie() {
	for i := 2; i < 100; i++ {
		//没开启协程
		//PrimeNum(i)
		//开启协程
		go PrimeNum(i)
	}
	//
	wg.Wait() // 阻塞主程序，直到所有协程完成
	fmt.Printf("共找到%v个素数", c)
}

// 接口
type sendmess struct {
	sendermes string
	type1     string
}

func (tm *sendmess) sendmesfun() {
	tm.sendermes = "这里是消息"
}
func (tm *sendmess) setype() {
	tm.type1 = "类型是消息"
}

type sendimg struct {
	senderimg string
	type1     string
}

func (im *sendimg) sendimgfun() {
	im.senderimg = "这里是图片"
}
func (im *sendimg) setype() {
	im.type1 = "类型是图片"
}

// 接口 ,接口定义了2个结构体的方法 ，接口的函数就是调用这个接口里定义的函数
type Mes interface {
	setype()
}

// 这里可以根据结构体来对应其他的方法调用
func Sendall(m Mes) {
	m.setype()
	switch expr11 := m.(type) {
	case *sendmess:
		expr11.sendmesfun()
	case *sendimg:
		expr11.sendimgfun()

	}
	fmt.Println("m=", m)
}

func Interface() {
	//赋值一个结构体
	tm := sendmess{}
	//
	Sendall(&tm)
	im := sendimg{}
	Sendall(&im)

}

// type struct
type User struct {
	name string
	id   int
}
type Contact struct {
	*User
	password string
}

func Sttu() {
	//var u1 User = User{
	//	name: "小李",
	//	id:   1,
	//}
	var u1 User
	//u1.name:="小李"
	//u1.id:=1
	//var u2 *Contact = &Contact{
	//	User: User{
	//		name: "小王",
	//		id:   2,
	//	},
	//	password: "123456",
	//}
	fmt.Printf("%s,%s", u1)
}

// 方法形参 User是个结构体，在函数名前面传入结构体形参，

func (u *User) printname() {
	fmt.Println(u.name)
}

func Method() {
	u := User{
		name: "结构体初始化name",
	}
	u.printname()
}

// map
func mapp() {
	m2 := map[string]string{
		"下午": "卖包子",
		"凌晨": "开车",
		"晚上": "滴滴车",
	}
	v, ok := m2["晚上"]
	if ok {
		fmt.Println("如果有值就打印value:", v)
	} else {
		fmt.Println("没有就打印这个哦")
	}
	delete(m2, "晚上")

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
