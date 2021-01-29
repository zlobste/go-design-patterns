package factory_method

import (
	"errors"
	"fmt"
)

const (
	CASH = iota
	CREDIT_CARD
)

type IPayment interface {
	Pay(amount float64) string
}

type Cash struct{}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("%#0.2f was paid in cash.\n", amount)
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("%#0.2f was paid by credit card.\n", amount)
}

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
