package handler

import (
	"net/http"
	"time"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type InboundOrderHandler struct {
	sv interfaces.IInboundOrderService
}

type InboundOrderJSON struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WareHouseID    int    `json:"warehouse_id"`
}

func NewInboundHandler(sv interfaces.IInboundOrderService) *InboundOrderHandler {
	return &InboundOrderHandler{
		sv: sv,
	}
}

func (h *InboundOrderHandler) PostInboundOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody InboundOrderJSON

	err := request.JSON(r, &reqBody)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the request body", nil))
		return
	}

	newInboundOrder := toInboundOrder(reqBody)

	entry, err := h.sv.Post(newInboundOrder)

	if err != nil {
		if err, ok := err.(*custom_error.InboundOrderErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("success", toInboundOrderJSON(entry)))
}

func toInboundOrder(inboundOrder InboundOrderJSON) model.InboundOrder {
	formatedDate, _ := time.Parse("2006-01-02", inboundOrder.OrderDate)

	return model.InboundOrder{
		ID:             inboundOrder.ID,
		OrderDate:      formatedDate,
		OrderNumber:    inboundOrder.OrderNumber,
		EmployeeID:     inboundOrder.EmployeeID,
		ProductBatchID: inboundOrder.ProductBatchID,
		WareHouseID:    inboundOrder.WareHouseID,
	}
}

func toInboundOrderJSON(inboundOrder model.InboundOrder) InboundOrderJSON {
	return InboundOrderJSON{
		ID:             inboundOrder.ID,
		OrderDate:      inboundOrder.OrderDate.Format("2006-01-02"),
		OrderNumber:    inboundOrder.OrderNumber,
		EmployeeID:     inboundOrder.EmployeeID,
		ProductBatchID: inboundOrder.ProductBatchID,
		WareHouseID:    inboundOrder.WareHouseID,
	}
}
