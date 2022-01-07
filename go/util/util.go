package util

import "fmt"

func SelectByKey(text ...string) (key int) {
	for i, v := range text {
		fmt.Printf("%v: %v\n", i+1, v)
	}
	fmt.Println("请选择:(数字)")
	fmt.Scanln(&key)
	return
}
