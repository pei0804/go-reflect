package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("----valueof-----")
	valueof(1, "str", false)
	fmt.Println("----valueof2-----")
	valueof2(100)
	fmt.Println("----valueof3-----")
	valueof3()
}

// vからValueOfでポインタにアクセスして、Elemでポインタにある値を返す
func valueof(value ...interface{}) {
	for _, v := range value {
		fmt.Println(reflect.ValueOf(&v).Elem())
	}
}

// valueに格納されている値がint型と適合したら、その値をansに格納する
func valueof2(value interface{}) {
	var ans int
	v := reflect.ValueOf(&value).Elem()
	if intValue, ok := v.Interface().(int); ok {
		ans += intValue
		fmt.Println(ans)
	}
}

// int型のゼロ値を取得し表示する
func valueof3() {
	var value int
	fmt.Println(reflect.ValueOf(&value).Elem())
}
