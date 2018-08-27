package entity

type Order struct {
	ID            int64  `json:"id"`
	InvoiceRefNum string `json:"invoice_ref_num"`
}
