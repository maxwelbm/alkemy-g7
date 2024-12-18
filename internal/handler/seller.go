package handler

import (
	"net/http"
	"strconv"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type SellersJSON struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func CreateHandlerSellers(service service.SellersService) *SellersController {
	return &SellersController{service: service}
}

type SellersController struct {
	service service.SellersService
}

func (hd *SellersController) GetAllSellers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// service
		sellers, err := hd.service.GetAll()
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request",
			})
			return
		}

		// response
		data := make(map[int]SellersJSON)
		for key, value := range sellers {
			data[key] = SellersJSON{
				ID:          value.ID,
				CID:         value.CID,
				CompanyName: value.CompanyName,
				Address:     value.Address,
				Telephone:   value.Telephone,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Get request executed successfully",
			"data":    data,
		})
	}
}

func (hd *SellersController) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request - Missing ID",
			})
			return
		}

		// service
		seller, err := hd.service.GetByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Get request executed successfully",
			"data":    seller,
		})
	}
}
