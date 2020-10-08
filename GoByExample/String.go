package main

import s "strings"
import "fmt"

// 我们给 `fmt.Println` 一个短名字的别名，我们随后将会经常用到。
var p = fmt.Println

func main() {

	// 这是一些 `strings` 中的函数例子。注意他们都是包中的
	// 函数，不是字符串对象自身的方法，这意味着我们需要考
	// 虑在调用时将字符串作为第一个参数进行传递。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))

	// 虽然不是 `strings` 的一部分，但是仍然值得一提的是获
	// 取字符串长度和通过索引获取一个字符的机制。
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
