package without

import (
	"log"
	"solid/solid/isp"
)

type OrderCanceled struct{}

func (oc *OrderCanceled) Next(o *isp.Order) {
	log.Println("NOT_IMPLEMENTED")
}

func (oc *OrderCanceled) Prev(o *isp.Order) {
	o.Status = "PENDING"
}

func (oc *OrderCanceled) Cancel(o *isp.Order) {
	log.Println("CANCELED")
}
