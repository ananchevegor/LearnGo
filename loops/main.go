package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("================")

	// это while просто в виде for
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("================")

	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("================")

	// ебать это ваще крутяк
	hi, there, _ := format(1, 2, 3, "Hi", "there", "!")

	fmt.Println(hi, there)

	fmt.Println(concat("hi", "go", "main"))

	for i := 0; i < 5; i++ {
		rec := recursion(i)
		fmt.Println(rec())
	}

	// функция будет сохранять предыдущее значение f
	// т. е. 100 - первое значение
	// 105 - второе значение (1*5)
	// 115 - третье значение : 105 сохранилось в f() + 2*5
	// 130 - четвертое значение : 115 сохранилось в f() + 3*5 и так далее
	f := beginNum(100)

	for i := 0; i < 5; i++ {
		fmt.Println(f(i * 5))
	}

}

// функции могут возвращать сразу несколько значений
func format(i, j, k int, a, b, c string) (string, string, int) {
	return a, b, i
}

// я могу явно указать какую именно переменную я хочу вернуть, крутяк!
func concat(a ...string) (result string) {
	result = strings.Join(a, ",")
	return
}

// я могу возвращать функции

func recursion(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func beginNum(begunNum int) func(i int) int {
	return func(i int) int {
		begunNum += i
		return begunNum
	}
}
