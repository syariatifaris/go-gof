package main

import "fmt"

type IPayment interface{
	MethodName()string
	Pay(name, email, phone string) error
}

type BankTransfer struct{}

func (b *BankTransfer) Transfer(){
	fmt.Println("transfer to bank")
}

type Paypal struct{}

func (*Paypal) MethodName() string{
	return "paypal"
}

func (p *Paypal) Pay(name, email, phone string) error {
	fmt.Println("pay with", p.MethodName())
	return nil
}

type BankTransferAdapter struct{
	*BankTransfer
}

func (*BankTransferAdapter) MethodName() string {
	return "bank transfer"
}

func (b *BankTransferAdapter) Pay(name, email, phone string) error {
	b.BankTransfer.Transfer()
	return nil
}

func NewPaymentProcess(method string) IPayment{
	switch method {
	case "bank_transfer":
		return &BankTransferAdapter{
			BankTransfer: new(BankTransfer),
		}
	case "paypal":
		return &Paypal{}
	}
	return nil
}

func main(){
	paypalPayment := NewPaymentProcess("paypal")
	paypalPayment.Pay("Faris", "test@test.com", "+123")

	bankTransfer := NewPaymentProcess("bank_transfer")
	bankTransfer.Pay("Faris", "test@test.com", "+123")
}

