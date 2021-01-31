## Decorator :package:

Decorator is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside special
wrapper objects.

### Description

Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same
interface. The resulting object will get a stacking behavior of all wrappers.

### Pros and Cons

| :white_check_mark: Pros  | :x: Cons                                               |
| --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| You can extend an object’s behavior without making a new subclass.                | It’s hard to remove a specific wrapper from the wrappers stack.                                                          |
| You can combine several behaviors by wrapping an object into multiple decorators. | It’s hard to implement a decorator in such a way that its behavior doesn’t depend on the order in the decorators stack.  |
| You can add or remove responsibilities from an object at runtime.                 | The initial configuration code of layers might look pretty ugly.                                                         |
| Single Responsibility Principle. You can divide a monolithic class that implements many possible variants of behavior into several smaller classes.

### Implementation

All object needed to implement the decorator pattern should implement same interface. Because of we must improve the
functionality of `GetPrice` method we'll create an interface like this:

```
type IComponent interface {
    GetPrice() int
}
```

And because of we are treating with a legacy recipe, all new ingredient must implement a `IComponent` interface.

```
type Pizza struct{}

func (p *Pizza) GetPrice() int {
	return PRICE_FOR_DOUGH
}
```

New ingredient:

```
type TomatoTopping struct {
    Pizza Pizza
}

func (c *TomatoTopping) GetPrice() int {
    pizzaPrice := c.Pizza.GetPrice()
    return pizzaPrice + PRICE_FOR_TOMATO
}
```

Another ingredient:

```
type CheeseTopping struct {
    Pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
    pizzaPrice := c.Pizza.GetPrice()
    return pizzaPrice + PRICE_FOR_CHEESE
}
```

### Important aspects:

* Unlike Adapter pattern, the object to be decorated is obtained by injection.
* Decorators should not alter the interface of an object.