package svcrepo

type Service struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Uom   string `json:"uom"`
}
