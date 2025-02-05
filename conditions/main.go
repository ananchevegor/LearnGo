package main

import (
	"fmt"
	"runtime"
)

func main() {

	k := 2
	if k == 0 {
		fmt.Println("k == 0")
	} else if k == 1 {
		fmt.Println(" K == 1")
	} else {
		fmt.Println("Other case")
	}

	// может быть пустой интерфейс и я могу задать туда переменную разного типа
	var val interface{}

	val = 9

	// val.(type) - проверка на то какой тип переменной
	// ок - это булевая переменная, ее принято называть ok

	// по сути тут проверяется ok если true, то Int, иначе Not Int
	// такая конструкция, нужно просто принять

	if iv, ok := val.(int); ok {
		fmt.Println("Int : ", iv)
	} else {
		fmt.Println("Not Int")
	}

	// По факту, os := sVal - присвоение переменной os занчения sVal и потом проврка уже os в switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X")
		fallthrough // fallthrough - продолжает switch даже после того, как одно из условий сработало, то есть по факту нет break

	case "linux":
		fmt.Println("Linux")

	default:
		fmt.Printf("%s.\n", os)
	}

	// до этого я создавал val - пустой интерфейс

	// проверка типа интерфейса

	switch tVal := val.(type) {
	case int32:
		fmt.Println(tVal)
	case bool:
		fmt.Println(tVal)

	default:
		fmt.Println("Default")

	}

}
