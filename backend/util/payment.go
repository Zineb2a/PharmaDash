package util

type PaymentInterface interface {
	ProcessPayment(total float32, cardNumber string, expiryDate string, cvv string)
	GetPaymentDetails(orderID int32)
}

type PaymentFactory struct {
}

type CreditPayment struct {
	OrderID    int32
	Total      float32
	CardNumber string
	ExpiryDate string
	cvv        string
}

type DebitPayment struct {
	OrderID    int32
	Total      float32
	CardNumber string
}

func (p *PaymentFactory) getPaymentStrategy(text string) (PaymentInterface, error) {
	if text == "credit" {
		return CreditPayment{}, nil
	}

	if text == "debit" {
		return DebitPayment{}, nil
	}
	return nil, nil
}

func (c CreditPayment) ProcessPayment(total float32, cardNumber string, expiryDate string, cvv string) {
}

func (d DebitPayment) ProcessPayment(total float32, cardNumber string, expiryDate string, cvv string) {
}

func (c CreditPayment) GetPaymentDetails(OrderID int32) {}

func (d DebitPayment) GetPaymentDetails(OrderID int32) {}
