package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
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

func (h *SectionController) GetAll(w http.ResponseWriter, r *http.Request) {
	s, err := h.sv.Get()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(err.Error(), nil))
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
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", data))
}

func (h *SectionController) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
	}

	s, err := h.sv.GetById(idInt)
	if err != nil {
		response.JSON(w, handleError(err), responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", s))
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

func handleError(err error) int {
	if errors.Is(err, repository.NotFoundError) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
