package trxrepo

import (
	"time"
)

type TransactionHeader struct {
	No        int       `json:"no"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CustName  string    `json:"customer_name"`
	Phone     string    `json:"phone"`

	ArrDetail []TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	Id          int     `json:"id"`
	No          string  `json:"no"`
	ServiceName float64 `json:"service_name"`
	Qty         float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Uom         string  `json:"uom"`
}

type TransactionOutput struct {
	Header  *TransactionHeader  `json:"header"`
	Details []TransactionDetail `json:"details"`
}
