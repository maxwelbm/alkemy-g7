package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
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

func (h *SectionController) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "invalid id param",
		})
	}

	s, err := h.sv.GetById(idInt)
	if err != nil {
		response.JSON(w, handleError(err), nil)
		return
	}

	response.JSON(w, http.StatusOK, s)
}

func (h *SectionController) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody SectionJSON
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "invalid request body",
		})
		return
	}

	section := model.Section{
		SectionNumber:      reqBody.SectionNumber,
		CurrentTemperature: reqBody.CurrentTemperature,
		MinimumTemperature: reqBody.MinimumTemperature,
		CurrentCapacity:    reqBody.CurrentCapacity,
		MinimumCapacity:    reqBody.MinimumCapacity,
		MaximumCapacity:    reqBody.MaximumCapacity,
		WarehouseID:        reqBody.WarehouseID,
		ProductTypeID:      reqBody.ProductTypeID,
	}

	s, err := h.sv.Post(section)
	if err != nil {
		response.JSON(w, handleError(err), nil)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]any{
		"message": "section created",
		"data":    s,
	})
}

func (h *SectionController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SectionController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "invalid id param",
		})
	}

	err = h.sv.Delete(idInt)
	if err != nil {
		response.JSON(w, handleError(err), nil)
		return
	}

	response.JSON(w, http.StatusNoContent, map[string]any{
		"message": "section removed successfully",
		"data":    nil,
	})
}

func handleError(err error) int {
	if errors.Is(err, repository.NotFoundError) {
		return http.StatusNotFound
	}
	if errors.Is(err, repository.ConflictError) {
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}
