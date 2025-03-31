package src

import "fmt"

// 在go语言中，权限都是通过首字母的大小来控制的
// Human父类
type Human struct {
	Name   string
	Age    int
	Gender string
}

// Human类绑定的方法Eat
func (p *Human) Eat() {
	fmt.Println(p.Name + "is Eating")
}

// student子类,嵌套Human
type Student struct {
	Hum    Human // 包含Human类里的变量,此时是类的嵌套
	School string
	Score  float64
}

// teacher类，去继承human
type Teacher struct {
	Human          // 直接写Human类型，没有字段名字
	Subject string //学科
}
