package sub

// 在go语言中，用一个层级不能有多个包
func Sub(a, b int) int {
	test5() //由于test5和sub.go在同一个包下面，所以可以直接使用，并且不需要sub.形式
	return a - b
}
