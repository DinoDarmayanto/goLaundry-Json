package trxrepo

import (
	"time"
)

type TransactionHeader struct {
	No        int
	StartDate time.Time
	EndDate   time.Time
	CustName  string
	Phone     string

	ArrDetail []TransactionDetail
}

type TransactionDetail struct {
	Id          int
	No          string
	ServiceName float64
	Qty         float64
	Price       float64
	Uom         string
}

type TransactionOutput struct {
	Header  *TransactionHeader  `json:"header"`
	Details []TransactionDetail `json:"details"`
}
