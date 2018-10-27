package ducks

type Duck struct {
	Types string
	Quantity int
	Strategy int
}

func (duck *Duck) SetStrategy(strateg int) {
	duck.Strategy = strateg
}

func (duck *Duck) DoIt(){
	if duck.Strategy == 0{
		duck.Fly()
	}

	if duck.Strategy == 1{
		duck.Dive()
	}

	if duck.Strategy == 2{
		duck.Quack()
	}

	if duck.Strategy == 3{
		duck.Swim()
	}
}

func (duck Duck) Fly() string{
	return "Я " + duck.Types + ", я Лечу"
}

func (duck Duck) Dive() string{
	return "Я " + duck.Types + ", я Ныряю"
}

func (duck Duck) Quack() string{
	return "Я " + duck.Types + ", я Крякая"
}

func (duck Duck) Swim() string{
	return "Я " + duck.Types + ", я Плыву"
}