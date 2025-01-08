package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
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

type ErrorResponse struct {
	Message string `json:"Message"`
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
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Unable to list Buyers", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyers))

}

func (bh *BuyerHandler) HandlerGetBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {

		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}

	buyer, err := bh.svc.GetBuyerByID(id)

	if err != nil {
		if errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
			response.JSON(w, http.StatusNotFound, ErrorResponse{
				Message: "Buyer Not Found",
			})
			return
		}

		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	responseBuyer := ResponseBuyerJson{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}

	data := Data{
		responseBuyer,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Erro ao serializar os dados",
		})
		return
	}

}

func (bh *BuyerHandler) HandlerDeleteBuyerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}

	err = bh.svc.DeleteBuyerByID(id)

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
		response.JSON(w, http.StatusNotFound, ErrorResponse{
			Message: "Buyer Not Found",
		})
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

		response.JSON(w, http.StatusUnprocessableEntity, nil)
		return
	}

	buyer, err := bh.svc.CreateBuyer(model.Buyer{
		CardNumberId: reqBody.CardNumberId,
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
	})

	if err != nil {
		if errors.Is(err.(*custom_error.CustomError).Err, custom_error.Conflict) {
			response.JSON(w, http.StatusConflict, ErrorResponse{
				Message: "card_number_id already exists",
			})
			return
		}
	}

	data := Data{
		Data: buyer,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Erro ao serializar os dados",
		})
		return
	}

}

func (bh *BuyerHandler) HandlerUpdateBuyer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}

	var reqBody RequestBuyerJson

	err = json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: "erro de desserialização dos dados",
		})
		return
	}

	buyer, err := bh.svc.UpdateBuyer(id, model.Buyer{
		CardNumberId: reqBody.CardNumberId,
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
	})

	if err != nil {
		var customErr custom_error.CustomError
		if errors.As(err, &customErr) {
			switch customErr.Err {
			case custom_error.NotFound:
				response.JSON(w, http.StatusNotFound, ErrorResponse{
					Message: "Buyer Not Found",
				})
				return
			case custom_error.EmptyFields:
				response.JSON(w, http.StatusUnprocessableEntity, ErrorResponse{
					Message: "At least one field must be mandatory to send the request",
				})
				return
			case custom_error.Conflict:
				response.JSON(w, http.StatusUnprocessableEntity, ErrorResponse{
					Message: "card_number_id already exists",
				})
				return
			}
		}

		response.JSON(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	data := Data{
		Data: buyer,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Erro ao serializar os dados",
		})
		return
	}

}
