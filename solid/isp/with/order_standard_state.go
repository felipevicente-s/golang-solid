package with

import "solid/solid/isp"

type OrderStandardState interface {
	Next(o *isp.Order)
	Prev(o *isp.Order)
}
