package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()
	fmt.Printf("1")

	duck2 := ToyDuck{} //就是一个toyduck实例
	duck2.Swim()
	fmt.Printf("2")

	// duck3 是 *ToyDuck=new(Todyduck) 就是指针nil ，
	duck3 := new(ToyDuck)
	duck3.Swim()
	fmt.Printf("3")

	// 当你声明这样的时候，Go 就帮你分配好内存 var duck4 ToyDuck
	// 不用担心空指针的问题，以为它压根就不是指针
	var duck4 ToyDuck //= duck4:=Toyduck{} or duck4=&toduck{}
	duck4.Swim()
	fmt.Printf("4")

	// duck5会分配一个指针，但是toduck并没有赋值，指针不知道指向哪里，所以这个就不对，
	//var duck5 *ToyDuck
	// 这边会直接panic 掉
	//duck5.Swim()

	// 赋值，初始化按字段名字赋值
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用
	duck7 := ToyDuck{"蓝色", 1024}
	duck7.Swim()

	// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "橘色"

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint64
}

// t toduck是接收器，要创建实例才能调用这个swim方法,如果前面没有接受实例就是一个包名直接调用就行，有接收器的必须要创建一个实例再调用方法
func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s，%d一只\n", t.Color, t.Price)
}
