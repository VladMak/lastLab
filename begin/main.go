package main

import (
	"./ducks"
	"fmt"
)

func main() {
	var input int
	fmt.Print("Введите количество Крякв: ")
	fmt.Scanln(&input)
	var kryakvs = ducks.Duck{Types: "Кряква", Quantity: input}

	fmt.Print("Введите количество Савок: ")
	fmt.Scanln(&input)
	var savks = ducks.Duck{Types:"Савка", Quantity: input}

	fmt.Print("Введите количество Нырков: ")
	fmt.Scanln(&input)
	var nirks = ducks.Duck{Types: "Нырок", Quantity: input}

	fmt.Print("Введите количество резиновых уточек: ")
	fmt.Scanln(&input)
	var rezins = ducks.Duck{Types: "Резиновая уточка", Quantity: input}

	fmt.Println("------------------")
	fmt.Println("Виды стратегий")
	fmt.Println("0 - лететь, 1 - нырять, 2 - крякать, 3 - плавать")
	fmt.Println("------------------")
	
	fmt.Print("Введите стратегию для Крякв: ")
	fmt.Scanln(&input)
	kryakvs.SetStrategy(input)

	fmt.Print("Введите стратегию для Савок: ")
	fmt.Scanln(&input)
	savks.SetStrategy(input)

	fmt.Print("Введите стратегию для Нырков: ")
	fmt.Scanln(&input)
	nirks.SetStrategy(input)

	fmt.Print("Введите стратегию для Резиновых уточек: ")
	fmt.Scanln(&input)
	rezins.SetStrategy(input)

	fmt.Println("--------------")
	fmt.Println(kryakvs.DoIt())
	fmt.Println(savks.DoIt())
	fmt.Println(nirks.DoIt())
	fmt.Println(rezins.DoIt())

	fmt.Println()
	fmt.Println("Для выхода нажмите любую клавишу...")
	fmt.Scanln()
}
