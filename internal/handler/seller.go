package handler

import (
	"strconv"
	"net/http"

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
	if ok := hd.handlerError(err, w); ok {
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
	if ok := hd.handlerError(err, w); ok {
		return
	}

	seller, err := hd.service.GetById(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller
	err := request.JSON(r, &seller)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	createdseller, err := hd.service.CreateSeller(&seller)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdseller))

}

func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	var s model.Seller
	err = request.JSON(r, &s)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	seller, err := hd.service.UpdateSeller(id, &s)
	if ok := hd.handlerError(err, w); ok {
		return
	}
		
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

func (hd *SellersController) DeleteSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	err = hd.service.DeleteSeller(id)
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
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Internal server error", nil))
		return true
	}
	return false
}
