package main

import "fmt"

type rect struct {
	width, height int
}

// 这里的 `area` 方法有一个_接收器(receiver)类型_ `rect`。
func (r *rect) area() int {
	fmt.Printf("r %+v\n", &r)
	r = &rect{width: 2, height: 2}
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。
// 这里是一个值类型接收器的例子。
func (r rect) perim() int {
	fmt.Printf("r %+v\n", &r)
	r.height = 0
	r.width = 0
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(&r)
	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Printf("%+v\n", r)
	fmt.Println("perim:", r.perim())
	fmt.Printf("%+v\n", r)

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，
	// 或者让方法能够改变接受的结构体。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Printf("%+v", r)
	fmt.Println("perim:", rp.perim())
	fmt.Printf("%+v", r)
}
