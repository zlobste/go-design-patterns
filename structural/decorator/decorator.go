package decorator

const (
	PRICE_FOR_DOUGH  = 10
	PRICE_FOR_TOMATO = 5
	PRICE_FOR_CHEESE = 10
)

type IPizza interface {
	GetPrice() int
}

type Pizza struct{}

func (p *Pizza) GetPrice() int {
	return PRICE_FOR_DOUGH
}

type TomatoTopping struct {
	Pizza Pizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + PRICE_FOR_TOMATO
}

type CheeseTopping struct {
	Pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + PRICE_FOR_CHEESE
}
