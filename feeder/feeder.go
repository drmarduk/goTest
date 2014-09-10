package feeder

import (
	"fmt"
)

type Feeder struct {
	Food string
}

func (f *Feeder) Feed(animal Animal) {
	animal.Eat(f.Food)
	fmt.Printf("%s\n", animal.Shit())
}
