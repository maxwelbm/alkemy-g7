package handler

import (
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

func (hd *SellersController) GetAllSellers(w http.ResponseWriter, r *http.Request) {
		sellers, err := hd.service.GetAll()
		if err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: err.Error()})
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
		response.JSON(w, http.StatusOK, model.ResponseBodySeller{Data: data})
}


func (hd *SellersController) GetById(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: "Missing int ID"})
			return
		}

		seller, err := hd.service.GetById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, model.ResponseBodyErrorSeller{Status: "Not Found", Message: err.Error()})
			return
		}

		response.JSON(w, http.StatusOK, model.ResponseBodySeller{Data: seller})

}

func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
		var seller model.Seller
		if err := request.JSON(r, &seller); err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: "Invalid JSON"})
			return
		}

		createdseller, err := hd.service.CreateSeller(seller)
		if err != nil {
			if err.Error() == "Seller's CID already exist" {
				response.JSON(w, http.StatusConflict, model.ResponseBodyErrorSeller{Status: "Conflict", Message: err.Error()})
				return
			} else {
				response.JSON(w, http.StatusUnprocessableEntity, model.ResponseBodyErrorSeller{Status: "Unprocessable Entity", Message: err.Error()})
				return
			}
		} 

		response.JSON(w, http.StatusCreated, model.ResponseBodySeller{Data: createdseller})

}

func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
		idSearch := chi.URLParam(r, "id")
        id, err := strconv.Atoi(idSearch)
        if err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: "Missing int ID"})
			return
		}

        if _, err := hd.service.GetById(id); err != nil {
            response.JSON(w, http.StatusNotFound, model.ResponseBodyErrorSeller{Status: "Not Found", Message: err.Error()})
            return
        }

        var s model.SellerUpdate
        if err := request.JSON(r, &s); err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: "Invalid JSON"})
			return
		}
        seller, err := hd.service.UpdateSeller(id, s)

        if err != nil {
			response.JSON(w, http.StatusBadRequest, model.ResponseBodyErrorSeller{Status: "Bad Request", Message: err.Error()})
            return
        } else {
			response.JSON(w, http.StatusOK, model.ResponseBodySeller{Data: seller})
        }
}
