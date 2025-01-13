package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IInboundOrderService interface {
	Post(inboundOrder model.InboundOrder) (model.InboundOrder, error)
}
