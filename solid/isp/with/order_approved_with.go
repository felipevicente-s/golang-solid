package with

import "solid/solid/isp"

type OrderApprovedWith struct{}

func (oaw *OrderApprovedWith) Next(o *isp.Order) {
	o.Status = "APPROVED"
}

func (oaw *OrderApprovedWith) Prev(o *isp.Order) {
	o.Status = "PENDING"
}
