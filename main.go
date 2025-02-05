package main

import (
	"fmt"
)

func main() {
	val := 9
	ptrVal := &val // знак & означает ссылку на ячейку памяти в которой хранится переменная val
	val = 20
	fmt.Println(*ptrVal)
}
