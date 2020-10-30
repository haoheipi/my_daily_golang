package cn

import "fmt"

//我们提供了一个类： 
//
// public class Foo {
//  public void first() { print("first"); }
//  public void second() { print("second"); }
//  public void third() { print("third"); }
//} 
//
// 三个不同的线程将会共用一个 Foo 实例。 
//
// 
// 线程 A 将会调用 first() 方法 
// 线程 B 将会调用 second() 方法 
// 线程 C 将会调用 third() 方法 
// 
//
// 请设计修改程序，以确保 second() 方法在 first() 方法之后被执行，third() 方法在 second() 方法之后被执行。 
//
// 
//
// 示例 1: 
//
// 输入: [1,2,3]
//输出: "firstsecondthird"
//解释: 
//有三个线程会被异步启动。
//输入 [1,2,3] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 second() 方法，线程 C 将会调用 third() 方法。
//正确的输出是 "firstsecondthird"。
// 
//
// 示例 2: 
//
// 输入: [1,3,2]
//输出: "firstsecondthird"
//解释: 
//输入 [1,3,2] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 third() 方法，线程 C 将会调用 second() 方法。
//正确的输出是 "firstsecondthird"。 
//
// 
//
// 提示： 
//
// 
// 尽管输入中的数字似乎暗示了顺序，但是我们并不保证线程在操作系统中的调度顺序。 
// 你看到的输入格式主要是为了确保测试的全面性。 
// 
// 👍 194 👎 0

// Time: 2020-10-24 13:40:29

//There is no code of Go type for this problem
var secondChan = make(chan int)
var thirdChan = make(chan int)
var mainChan = make(chan int)

func first() {
	fmt.Print("first")
	secondChan <- 1
}

func second() {
	signal := <-secondChan
	fmt.Print("second")
	thirdChan <- signal
	close(secondChan)
}

func third() {
	signal := <-thirdChan
	fmt.Print("third")
	mainChan <- signal
	close(thirdChan)
}

func main() {
	funcMap := map[int]func(){1: first, 2: second, 3: third}
	inputList := [3]int{1, 2, 3}

	for _, num := range inputList {
		go funcMap[num]()
	}

	_ = <-mainChan
	close(mainChan)
}
