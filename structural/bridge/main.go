package main

import "fmt"

type IPaymentMethod interface {
	Pay(username, email string)
}

type bankTransfer struct{}

func (b *bankTransfer) Pay(username, email string) {
	fmt.Println("paid using bank transfer", email, username)
}

type creditCard struct{}

func (*creditCard) Pay(username, email string) {
	fmt.Println("paid using credit card", email, username)
}

type vcc struct{}

func (*vcc) Pay(username, email string) {
	fmt.Println("paid using vcc", email, username)
}

//bridge
type Payment struct {
	PaymentMethod IPaymentMethod
	Username, Email string
}

func newTransaction(txType, username, email string) ITransaction{
	switch txType {
	case "marketplace":
		return &marketplaceTransaction{
			Payment: &Payment{
				Username: username,
				Email: email,
			},
		}
	case "digital":
		return &digitalTransaction{
			Payment: &Payment{
				Username: username,
				Email: email,
			},
		}
	}
	return nil
}

type ITransaction interface{
	DoPay()
	SetPaymentMethod(method IPaymentMethod)
}

type marketplaceTransaction struct {
	*Payment
}

func (m *marketplaceTransaction) SetPaymentMethod(method IPaymentMethod) {
	m.PaymentMethod = method
}

func (m *marketplaceTransaction) DoPay() {
	m.PaymentMethod.Pay(m.Username, m.Email)
}

type digitalTransaction struct {
	*Payment
}

func (d *digitalTransaction) DoPay() {
	d.PaymentMethod.Pay(d.Username, d.Email)
}

func (d *digitalTransaction) SetPaymentMethod(method IPaymentMethod) {
	d.PaymentMethod = method
}

func main(){
	vccMethod := new(vcc)
	bankTransfer := new(bankTransfer)
	ccMethod := new(creditCard)

	fmt.Println("--marketplace transaction--")
	marketplaceTx := newTransaction("marketplace", "faris", "faris.syariati@tokopedia.com")
	marketplaceTx.SetPaymentMethod(vccMethod)
	marketplaceTx.DoPay()
	marketplaceTx.SetPaymentMethod(bankTransfer)
	marketplaceTx.DoPay()

	fmt.Println("--digital transaction--")
	digitalTx := newTransaction("digital", "fathiya", "fathiya@tokopedia.com")
	digitalTx.SetPaymentMethod(ccMethod)
	digitalTx.DoPay()
}





