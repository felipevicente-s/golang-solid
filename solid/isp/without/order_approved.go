package without

import (
	"log"
	"solid/solid/isp"
)

type OrderApproved struct{}

func (oa *OrderApproved) Next(o *isp.Order) {
	o.Status = "APPROVED"
}

func (oa *OrderApproved) Prev(o *isp.Order) {
	o.Status = "PENDING"
}

func (oa *OrderApproved) Cancel(o *isp.Order) {
	log.Println("NOT_IMPLEMENTED")
}
