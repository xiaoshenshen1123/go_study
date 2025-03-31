package main

// 在go语言中没有枚举类型，但是我们可以使用const+iota（常量累加器）来进行模拟
// 模拟表示一周
const (
	MONDAY = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
	m, n = iota, iota
)
const (
	JANU = iota
)

/*
1.iota是常量组计数器
2.iota从0开始，每换行递增1
3.常量组有个特点：如果不赋值那么默认与上一行表达式相同
4.如果同一行出现两个iota，那么两个iota的值是相同的
5.每个常量组的iota是独立的，如果遇到新的const常量组 iota会清零
*/

func main() {
	//var number int
	//var name string
	//var flag bool
	// 可以使用变量组来统一定义变量
	//var (
	//	number int
	//	name   stirng
	//	flag   bool
	//)
}
