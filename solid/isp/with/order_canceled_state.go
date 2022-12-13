package with

import "solid/solid/isp"

type OrderCanceledState interface {
	Prev(o *isp.Order)
	Cancel(o *isp.Order)
}
