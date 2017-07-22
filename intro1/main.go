package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("----typeof-----")
	typeof(1, "string", true)
	fmt.Println("----typeof2-----")
	typeof2(1)
}

// vが何の値かを出力する
func typeof(value ...interface{}) {
	for _, v := range value {
		fmt.Println(reflect.TypeOf(v))
	}
}

// vがint型と同じか比較する
func typeof2(v interface{}) {
	fmt.Println(reflect.TypeOf(v).Kind() == reflect.Int)
}
