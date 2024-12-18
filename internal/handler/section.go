package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type SectionsJSON struct {
	Sections []SectionJSON `json:"sections"`
}

type SectionJSON struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

func CreateHandlerSections(sv *service.SectionService) *SectionController {
	return &SectionController{sv}
}

type SectionController struct {
	sv *service.SectionService
}

func (h *SectionController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := h.sv.Get()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := make(map[int]SectionJSON)
		for key, value := range s {
			data[key] = SectionJSON{
				ID:                 value.ID,
				SectionNumber:      value.SectionNumber,
				CurrentTemperature: value.CurrentTemperature,
				MinimumTemperature: value.MinimumTemperature,
				CurrentCapacity:    value.CurrentCapacity,
				MinimumCapacity:    value.MinimumCapacity,
				MaximumCapacity:    value.MaximumCapacity,
				WarehouseID:        value.WarehouseID,
				ProductTypeID:      value.ProductTypeID,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *SectionController) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SectionController) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SectionController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SectionController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
