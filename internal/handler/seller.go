package handler

import (
	"errors"
	"net/http"
	"strconv"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type SellersJSON struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func CreateHandlerSellers(service interfaces.ISellerService) *SellersController {
	return &SellersController{service: service}
}

type SellersController struct {
	service interfaces.ISellerService
}

func (hd *SellersController) createJSONReturnError(status string, message string) *model.ResponseBodyErrorSeller {
	return &model.ResponseBodyErrorSeller{Status: status, Message: message}
}

func (hd *SellersController) createJSONReturn(data any) *model.ResponseBodySeller {
	return &model.ResponseBodySeller{Data: data}
}

func (hd *SellersController) GetAllSellers(w http.ResponseWriter, r *http.Request) {
		sellers, err := hd.service.GetAll()
		if err != nil {
			response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, err.Error()))
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
			})
		}
		response.JSON(w, http.StatusOK, hd.createJSONReturn(data))
}


func (hd *SellersController) GetById(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, model.ErrorMissingID.Error()))
			return
		}

		seller, err := hd.service.GetById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, hd.createJSONReturnError(model.StatusNotFound, err.Error()))
			return
		}

		response.JSON(w, http.StatusOK, hd.createJSONReturn(seller))

}

func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
		var seller model.Seller
		if err := request.JSON(r, &seller); err != nil {
			response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, model.ErrorInvalidJSON.Error()))
			return
		}

		createdseller, err := hd.service.CreateSeller(seller)
		if err != nil {
			if ok := errors.Is(err, model.ErrorCIDAlreadyExist); ok {
				response.JSON(w, http.StatusConflict, hd.createJSONReturnError(model.StatusConflict, err.Error()))
				return
			} else {
				response.JSON(w, http.StatusUnprocessableEntity, hd.createJSONReturnError(model.StatusUnprocessableEntity, err.Error()))
				return
			}
		} 

		response.JSON(w, http.StatusCreated, hd.createJSONReturn(createdseller))

}

func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
		idSearch := chi.URLParam(r, "id")
        id, err := strconv.Atoi(idSearch)
        if err != nil {
			response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, model.ErrorMissingID.Error()))
			return
		}

        if _, err := hd.service.GetById(id); err != nil {
			response.JSON(w, http.StatusNotFound, hd.createJSONReturnError(model.StatusNotFound, err.Error()))
            return
        }

        var s model.SellerUpdate
        if err := request.JSON(r, &s); err != nil {
			response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, model.ErrorInvalidJSON.Error()))
			return
		}
        seller, err := hd.service.UpdateSeller(id, s)

        if err != nil {
			if ok := errors.Is(err, model.ErrorCIDAlreadyExist); ok {
				response.JSON(w, http.StatusConflict, hd.createJSONReturnError(model.StatusConflict, err.Error()))
				return
			} else {
				response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, err.Error()))
				return
			}
        } else {
			response.JSON(w, http.StatusOK, hd.createJSONReturn(seller))
        }
}

func (hd *SellersController) DeleteSellers(w http.ResponseWriter, r *http.Request) {
	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, model.ErrorMissingID.Error()))
		return
	}

	if _, err := hd.service.GetById(id); err != nil {
		response.JSON(w, http.StatusNotFound, hd.createJSONReturnError(model.StatusNotFound, err.Error()))
		return
	}

	err = hd.service.DeleteSeller(id)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, hd.createJSONReturnError(model.StatusBadRequest, err.Error()))
		return
	} else {
		response.JSON(w, http.StatusNoContent, hd.createJSONReturn(""))
	}
}
