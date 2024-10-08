package main

import "fmt"

func main() {
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法
	// 这一句会编译错误
	//fake.Swim()
	fake.FakeSwim()

	// 转换为Fish
	td := Fish(fake)
	// 真的变成了鱼
	td.Swim()

	sFake := StrongFakeFish{}
	// 这里就是调用了自己的方法
	sFake.Swim()

	td = Fish(sFake)
	// 真的变成了鱼
	td.Swim()

	vividug := dugtru{
		Color:  "黄色的毛毛",
		mouth1: "扁扁的嘴巴",
	}
	vividug.Duggg()

}

type dugtru struct {
	Color  string
	mouth1 string
}

func (d *dugtru) Duggg() {
	fmt.Println("鸭子的结构体：", d.Color, d.mouth1)
}

// 定义了一个新类型，注意是新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎\n")
}

// 定义了一个新类型
type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎\n")
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是一直鸭子\n")
}
