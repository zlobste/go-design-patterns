## Factory method :factory:

### Description

The Factory Method pattern is a design pattern used to define a runtime interface for creating an object. It’s called a
factory because it creates various types of objects without necessarily knowing what kind of object it creates or how to
create it.

Allows the sub-classes to choose the type of objects to create at runtime It provides a simple way of extending the
family of objects with minor changes in application code. Promotes the loose-coupling by eliminating the need to bind
application-specific structs into the code.

### Pros and Cons

| :white_check_mark: Pros                                                 | :x: Cons                                                                                                                                                                                                                        |
| ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| You avoid tight coupling between the creator and the concrete products. | The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern. The best case scenario is when you’re introducing the pattern into an existing hierarchy of creator classes. |
| Single Responsibility Principle. You can move the product creation code into one place in the program, making the code easier to support.
| Open/Closed Principle. You can introduce new types of products into the program without breaking existing client code.

### Implementation

The factory method defines the interface for creating objects but allows subclasses to decide which classes to create.
In this example, we will use a template to create a payment object model.

You can pay for purchases in different ways. In this example, we will present only two of them (in cash and by credit
card). In the context of the Factory Method design pattern, they are our product.

Each form of payment implements an interface that provides a payment function:

```
type IPayment interface {
    Pay(amount float64) string
}
```

Now you can create any number of payment types with the same functions (implement the same interface):

```
type Cash struct{}

func (c *Cash) Pay(amount float64) string {
    return fmt.Sprintf("%#0.2f was paid in cash.\n", amount)
}
```

```
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("%#0.2f was paid by credit card.\n", amount)
}
```

An additional thing you would require is a factory that would return different payment types (in cash/by credit card)
based on the user's request:

```
func GetPaymentMethod(method int64) (IPayment, error) {
	switch method {
	case CASH:
		return new(Cash), nil
	case CREDIT_CARD:
		return new(CreditCard), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d is not recognized.\n", method))
	}
}
```

### Important aspects when we implement the Factory Method design pattern are:

* Designing the arguments of the factory method.
* Considering an internal object pool that will allow object cache and reuse instead of created from scratch.