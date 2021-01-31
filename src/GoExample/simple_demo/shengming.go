package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ts1 []int
	ts2 := make([]int, 0, 12)
	ts3 := new([]int)
	fmt.Println(reflect.DeepEqual(ts1, ts2))
	fmt.Println(reflect.DeepEqual(ts2, *ts3))
	fmt.Println(reflect.DeepEqual(ts1, *ts3))
	ts1 = append(ts1, 1)
	ts2 = append(ts2, 1)
	*ts3 = append(*ts3, 1)
	fmt.Println(reflect.DeepEqual(ts1, ts2))
	fmt.Println(reflect.DeepEqual(ts2, *ts3))
	fmt.Println(reflect.DeepEqual(ts1, *ts3))
}
