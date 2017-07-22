package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("sample.txt")
	HandleData1(f)
	HandleData2("a")
}

func HandleData1(x interface{}) {
	f, ok := x.(*os.File)
	if ok {
		fmt.Println("file is " + f.Name())
	}
}

func HandleData2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("intパターン")
	case string:
		fmt.Println("stringパターン")
	case map[string]int:
		fmt.Println("map[string]intパターン")
	default:
		fmt.Println("???")
	}
}
