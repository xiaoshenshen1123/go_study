package main

import "fmt"

//实现go多态，需要定义接口
//人类的武器发起攻击，不同登记子弹效果不同

// 定义一个接口，注意类型是interface
type IAttack interface {
	//接口函数可以有多个，但是只能有函数原型，不可以实现
	Attack()
}

// 低等级
type HumanLowLevel struct {
	name  string
	level int
}

// 高等级
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ",造成1000点伤害")
}
func (a *HumanHighLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ",造成5000点伤害")
}

// 定义一个多态的通用接口,传入不同的对象，调用一样的方法实现不同的效果==》多态
func DoAttack(a IAttack) {
	a.Attack()
}
func main() {
	//var player interface{}
	//定义一个包含Attack的接口变量
	var player IAttack
	lowLevel := HumanLowLevel{
		name:  "小明",
		level: 1,
	}
	highLevel := HumanHighLevel{
		name:  "占山",
		level: 10,
	}
	lowLevel.Attack()
	highLevel.Attack()

	//对player赋值为lowlevel,接口需要使用指针类型赋值
	player = &lowLevel
	player.Attack()
	//对player赋值为highlevel,接口需要使用指针类型赋值
	player = &highLevel
	player.Attack()
	fmt.Println("----------多态")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
