package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type BuyerHandler struct {
	svc *service.BuyerService
}

type ResponseBuyerJson struct {
	Id           int    `json:"id"`
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Data struct {
	Data []ResponseBuyerJson `json:"data"`
}

func NewBuyerHandler(svc *service.BuyerService) *BuyerHandler {
	return &BuyerHandler{svc}
}

func (bh *BuyerHandler) HandlerGetAllBuyers(w http.ResponseWriter, r *http.Request) {

	buyers, err := bh.svc.GetAllBuyer()
	if err != nil {

		response.JSON(w, http.StatusNotFound, map[string]any{
			"message": "Não há buyers cadastrados",
		})
		return
	}

	var dataBuyers []ResponseBuyerJson
	for _, b := range buyers {

		dataBuyers = append(dataBuyers, ResponseBuyerJson{
			Id:           b.Id,
			CardNumberId: b.CardNumberId,
			FirstName:    b.FirstName,
			LastName:     b.LastName,
		})
	}

	data := Data{
		Data: dataBuyers,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {

		http.Error(w, "Erro ao serializar os dados", http.StatusInternalServerError)
		return
	}
}
