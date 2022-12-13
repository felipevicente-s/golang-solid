package without

import "solid/solid/isp"

type OrderState interface {
	Next(o *isp.Order)
	Prev(o *isp.Order)
	Cancel(o *isp.Order)
}
