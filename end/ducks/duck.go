package ducks

type Duck struct {
	Types string
	Quantity int
	Strategy int
}

func (duck *Duck) SetStrategy(strateg int) {
	duck.Strategy = strateg
}

func (duck *Duck) DoIt() string{
    switch duck.Strategy{
        case 0:
            return duck.Fly()
        case 1:
            return duck.Dive()
        case 2:
            return duck.Quack()
        case 3:
            return duck.Swim()
    }
    return ""
}

func (duck Duck) Fly() string{
	return "Я " + duck.Types + ", я Лечу"
}

func (duck Duck) Dive() string{
	return "Я " + duck.Types + ", я Ныряю"
}

func (duck Duck) Quack() string{
	return "Я " + duck.Types + ", я Крякаю"
}

func (duck Duck) Swim() string{
	return "Я " + duck.Types + ", я Плыву"
}
