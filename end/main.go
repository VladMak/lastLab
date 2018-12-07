package main

import (
	"./ducks"
	"fmt"
)

func main() {
	var input int
    ducksNames := []string{"Кряква", "Савка", "Нырок", "Резиновая уточка"}
    var allDucks []ducks.Duck
    for _, v := range ducksNames{
        fmt.Print("Введите количество уток семейства (" + v + "): ")
        fmt.Scanln(&input)
        allDucks = append(allDucks, ducks.Duck{Types: v, Quantity: input})
    }

	fmt.Println("------------------")
	fmt.Println("Виды стратегий")
	fmt.Println("0 - лететь, 1 - нырять, 2 - крякать, 3 - плавать")
	fmt.Println("------------------")

    for i := 0; i < len(allDucks); i++{
        fmt.Print("Введите стратегию для " + allDucks[i].Types + ": ")
        fmt.Scanln(&input)
        allDucks[i].SetStrategy(input)
    }

    fmt.Println("--------------")
    for _, v := range allDucks {
	    fmt.Println(v.DoIt())
    }
	fmt.Println("--------------")

	fmt.Println()
	fmt.Println("Для выхода нажмите любую клавишу...")
	fmt.Scanln()
}
