package main

import "reflect"

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)
}
