package main

import (
	"fmt"
)

func NewTransactiontion(payment IPayment, total int, email string) *Transaction {
	return &Transaction{
		payment: payment,
		total:   total,
		email:   email,
	}
}

type Transaction struct {
	payment IPayment
	email   string
	total   int
}

func (t *Transaction) DoPayment() {
	t.payment.Pay(t.total, t.email)
}

func (t *Transaction) SetStrategy(payment IPayment) {
	t.payment = payment
}

type IPayment interface {
	Pay(total int, email string)
}

type BankTransfer struct{}

func (*BankTransfer) Pay(total int, email string) {
	fmt.Println("pay with bank transfer, Rp:", total, "email:", email)
}

type CreditCard struct{}

func (*CreditCard) Pay(total int, email string) {
	fmt.Println("pay with credit card installment, Rp:", total, "email:", email)
}

func main() {
	bankStrategy := new(BankTransfer)
	ccStrategy := new(CreditCard)

	transaction := NewTransactiontion(bankStrategy, 10000, "syariati.faris@gmail.com")
	transaction.DoPayment()

	transaction.SetStrategy(ccStrategy)
	transaction.DoPayment()
}
