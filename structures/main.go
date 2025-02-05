package main

import "fmt"

type Human struct {
	Age    uint32
	Name   string
	Weight *Params // можно сделать ссылку на структуру - это означает, что структура Human будет весить маньше, потому что Params нет в ячейке этой структуры
}

type Human2 struct {
	Age    uint32
	Name   string
	Weight Params // можно не делать ссылку на структуру - это означает, что структура Human будет весить больше, потому что Params в ячейке этой структуры
}

type Human3 struct {
	Age    uint32
	Name   string
	Params // встроенная структура Params
}
type Human4 struct {
	Age    uint32
	Name   string
	weight float32 // можно обращаться к двум weight, главное явно указывать к какому
	Params         // встроенная структура Params
}

type Params struct {
	weight float32
}

type GoVersion string

const (
	GoVersion120 GoVersion = "1.20"
	GoVersion119 GoVersion = "1.19"
)

func main() {
	// присвоение структур

	human := Human{
		Age:  23,
		Name: "Egor",
		Weight: &Params{
			weight: 57,
		},
	}

	var (
		alex   Human2
		kirill Human3
		human2 *Human2
		max    Human4
	)

	fmt.Println(alex, kirill, human2, max)

	human2 = new(Human2) // можно создать структуру с возвращением указателя на эту структуру
	fmt.Println(human2)

	fmt.Println(human.Name)

	fmt.Println(GoVersion120) // константами я могу пользоваться везде, главное, чтобы она была определена в файле

}
