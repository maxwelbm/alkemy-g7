package handler

import (
	"net/http"
	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type SellersJSON struct {
	ID          int    		`json:"id"`
	CID         int 		`json:"cid"`
	CompanyName string    	`json:"company_name"`
	Address     string 		`json:"address"`
	Telephone   string  	`json:"telephone"`
}

func CreateHandlerSellers(service interfaces.ISellerService) *SellersController {
	return &SellersController{service: service}
}

type SellersController struct {
	service interfaces.ISellerService
}

func (hd *SellersController) GetAllSellers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// service
		sellers, err := hd.service.GetAll()
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request",
				"data":    nil,
			})
			return
		}

		// response
		data := make(map[int]SellersJSON)
		for key, value := range sellers {
			data[key] = SellersJSON{
				ID:          	value.ID,
				CID:        	value.CID,
				CompanyName:    value.CompanyName,
				Address:  		value.Address,
				Telephone:   	value.Telephone,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Get request executed successfully",
			"data":    data,
		})
	}
}