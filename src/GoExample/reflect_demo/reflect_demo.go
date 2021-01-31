package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	rv := []interface{}{"hi", 1, func() {}, &sync.Mutex{}}
	for _, v := range rv {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s\n", v.Kind())
		}
	}

	blog := Blog{"煎鱼"}
	typeof := reflect.TypeOf(blog)
	fmt.Println(typeof.String())

	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x))
}

type Blog struct {
	name string
}
