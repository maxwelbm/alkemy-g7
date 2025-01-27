package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

func CreateHandlerSellers(service interfaces.ISellerService) *SellersController {
	return &SellersController{Service: service}
}

type SellersController struct {
	Service interfaces.ISellerService
}

func (hd *SellersController) GetAllSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := hd.Service.GetAll()
	if ok := hd.handlerError(err, w); ok {
		return
	}

	data := make([]model.Seller, 0)
	for _, value := range sellers {
		data = append(data, model.Seller{
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

func (hd *SellersController) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id == 0 || err != nil {
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	if ok := hd.handlerError(err, w); ok {
		return
	}

	seller, err := hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller
	if err := request.JSON(r, &seller); err != nil {
		response.JSON(w, er.ErrInvalidSellerJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidSellerJSONFormat.Error(), nil))
		return
	}

	createdseller, err := hd.Service.CreateSeller(&seller)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdseller))
}

func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)

	if id == 0 || err != nil {
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	if ok := hd.handlerError(err, w); ok {
		return
	}

	_, err = hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	var s model.Seller
	if err := request.JSON(r, &s); err != nil {
		response.JSON(w, er.ErrInvalidSellerJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidSellerJSONFormat.Error(), nil))
		return
	}

	seller, err := hd.Service.UpdateSeller(id, &s)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

func (hd *SellersController) DeleteSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)

	if id == 0 || err != nil {
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	if ok := hd.handlerError(err, w); ok {
		return
	}

	_, err = hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	err = hd.Service.DeleteSeller(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("", nil))
}

func (hd *SellersController) handlerError(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err, ok := err.(*er.SellerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return true
		}

		if err, ok := err.(*er.LocalityError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return true
		}
		
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unmapped seller handler error", nil))

		return true
	}

	return false
}
