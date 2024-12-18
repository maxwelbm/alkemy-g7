package handler

import (
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type BuyerHandler struct {
	svc *service.BuyerService
}

func NewBuyerHandler(svc *service.BuyerService) *BuyerHandler {
	return &BuyerHandler{svc}
}

func (bh *BuyerHandler) HandlerGetAllBuyers(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Bateu aqui")
	buyers, err := bh.svc.GetAllBuyer()

	

	if err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "Não há buyers cadastrados",
		})
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "Success",
		"Data":    buyers,
	})
}
