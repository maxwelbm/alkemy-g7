package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type BuyerHandler struct {
	svc interfaces.IBuyerservice
}

type ResponseBuyerJson struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type RequestBuyerJson struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Data struct {
	Data any
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

func (bh *BuyerHandler) HandlerGetBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	buyer, err := bh.svc.GetBuyerByID(id)

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	response := ResponseBuyerJson{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}

	data := Data{
		response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {

		http.Error(w, "Erro ao serializar os dados", http.StatusInternalServerError)
		return
	}

}

func (bh *BuyerHandler) HandlerDeleteBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = bh.svc.DeleteBuyerByID(id)

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	w.Write([]byte(""))

}

func (bh *BuyerHandler) HandlerCreateBuyer(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBuyerJson

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil || reqBody.CardNumberId == "" || reqBody.FirstName == "" || reqBody.LastName == "" {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	buyer, err := bh.svc.CreateBuyer(model.Buyer{
		CardNumberId: reqBody.CardNumberId,
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
	})

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.Conflict) {
		http.Error(w, "card_number_id already exists", http.StatusConflict)
		return
	}

	data := Data{
		Data: buyer,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {

		http.Error(w, "Erro ao serializar os dados", http.StatusInternalServerError)
		return
	}

}
