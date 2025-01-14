package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type SellersJSON struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	Locality    int    `json:"locality_id"`
}

func CreateHandlerSellers(service interfaces.ISellerService) *SellersController {
	return &SellersController{service: service}
}

type SellersController struct {
	service interfaces.ISellerService
}

func (hd *SellersController) GetAllSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := hd.service.GetAll()
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	data := make([]SellersJSON, 0)
	for _, value := range sellers {
		data = append(data, SellersJSON{
			ID:          value.ID,
			CID:         value.CID,
			CompanyName: value.CompanyName,
			Address:     value.Address,
			Telephone:   value.Telephone,
			Locality:    value.Locality,
		})
	}
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", data))
}

func (hd *SellersController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrorMissingSellerID.Error(), nil))
		return
	}

	seller, err := hd.service.GetById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller
	if err := request.JSON(r, &seller); err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrorInvalidSellerJSONFormat.Error(), nil))
		return
	}

	createdseller, err := hd.service.CreateSeller(&seller)
	if err != nil {
		if ok := errors.Is(err, er.ErrorCIDSellerAlreadyExist); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else if ok := errors.Is(err, er.ErrorLocalityNotFound); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else if ok := errors.Is(err, model.ErrorLocalityNotFound); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdseller))

}

func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrorMissingSellerID.Error(), nil))
		return
	}

	var s model.Seller
	if err := request.JSON(r, &s); err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrorInvalidSellerJSONFormat.Error(), nil))
		return
	}

	seller, err := hd.service.UpdateSeller(id, &s)

	if err != nil {
		if ok := errors.Is(err, er.ErrorCIDSellerAlreadyExist); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else if ok := errors.Is(err, er.ErrorLocalityNotFound); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else if ok := errors.Is(err, model.ErrorLocalityNotFound); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}
	} else {
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
	}
}

func (hd *SellersController) DeleteSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrorMissingSellerID.Error(), nil))
		return
	}

	err = hd.service.DeleteSeller(id)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(err.Error(), nil))
		return
	} else {
		response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("", nil))
	}
}
