package main

import (
	"./ducks"
	"fmt"
)

func main() {
	var input int
	fmt.Println("Введите количество Крякв")
	fmt.Scanln(&input)
	var kryakvs = ducks.Duck{Types: "Кряква", Quantity: input}

	fmt.Println("Введите количество Савок")
	fmt.Scanln(&input)
	var savks = ducks.Duck{Types:"Савка", Quantity: input}

	fmt.Println("Введите количество Нырков")
	fmt.Scanln(&input)
	var nirks = ducks.Duck{Types: "Нырок", Quantity: input}

	fmt.Println("Введите количество резиновых уточек")
	fmt.Scanln(&input)
	var rezins = ducks.Duck{Types: "Резиновая уточка", Quantity: input}

	fmt.Println()
	fmt.Println("Виды стратегий")
	fmt.Println("0 - лететь, 1 - нырять, 2 - крякать, 3 - плавать")
	
	fmt.Println("Введите стратегию для Крякв")
	fmt.Scanln(&input)
	kryakvs.SetStrategy(input)

	fmt.Println("Введите стратегию для Савок")
	fmt.Scanln(&input)
	savks.SetStrategy(input)

	fmt.Println("Введите стратегию для Нырков")
	fmt.Scanln(&input)
	nirks.SetStrategy(input)

	fmt.Println("Введите стратегию для Резиновых уточек")
	fmt.Scanln(&input)
	rezins.SetStrategy(input)

	kryakvs.DoIt()
	savks.DoIt()
	nirks.DoIt()
	rezins.DoIt()
}