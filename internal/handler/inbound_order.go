package handler

import (
	"net/http"
	"time"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type InboundOrderHandler struct {
	sv  interfaces.IInboundOrderService
	log logger.Logger
}

type InboundOrderJSON struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WareHouseID    int    `json:"warehouse_id"`
}

func NewInboundHandler(sv interfaces.IInboundOrderService, log logger.Logger) *InboundOrderHandler {
	return &InboundOrderHandler{
		sv:  sv,
		log: log,
	}
}

func (h *InboundOrderHandler) PostInboundOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Log("InboundOrderHandler", "INFO", "Received request to create inbound order")

	var reqBody InboundOrderJSON

	err := request.JSON(r, &reqBody)

	if err != nil {
		h.log.Log("InboundOrderHandler", "ERROR", "Error parsing request body: "+err.Error())
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the request body", nil))

		return
	}

	newInboundOrder := toInboundOrder(reqBody)

	entry, err := h.sv.Post(newInboundOrder)

	if err != nil {
		if err, ok := err.(*customerror.InboundOrderErr); ok {
			h.log.Log("InboundOrderHandler", "ERROR", "Business error: "+err.Error())
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))

			return
		}

		h.log.Log("InboundOrderHandler", "ERROR", "Internal server error: "+err.Error())
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	h.log.Log("InboundOrderHandler", "INFO", "Inbound order created successfully")
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
