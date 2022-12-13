package with

import "solid/solid/isp"

type OrderFailedWith struct{}

func (ofw *OrderFailedWith) Prev(o *isp.Order) {
	o.Status = "PENDING"
}

func (ofw *OrderFailedWith) Cancel(o *isp.Order) {
	o.Status = "CANCELED"
}
