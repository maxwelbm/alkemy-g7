package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

type SectionsJSON struct {
	Sections []SectionJSON `json:"sections"`
}

type SectionJSON struct {
	ID                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
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

	data := []SectionJSON{}
	for _, value := range s {
		data = append(data, SectionJSON{
			ID:                 value.ID,
			SectionNumber:      value.SectionNumber,
			CurrentTemperature: value.CurrentTemperature,
			MinimumTemperature: value.MinimumTemperature,
			CurrentCapacity:    value.CurrentCapacity,
			MinimumCapacity:    value.MinimumCapacity,
			MaximumCapacity:    value.MaximumCapacity,
			WarehouseID:        value.WarehouseID,
			ProductTypeID:      value.ProductTypeID,
		})

	}
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", data))
}

func (h *SectionController) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		return
	}

	s, err := h.sv.GetById(idInt)
	if err != nil {
		response.JSON(w, handleError(err), responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", s))
}

func (h *SectionController) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody SectionJSON

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid request body", nil))
		return
	}

	if reqBody == (SectionJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
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
		response.JSON(w, handleError(err), responses.CreateResponseBody(err.Error(), nil))
		return
	}
	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("section created", s))
}

func (h *SectionController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		return
	}

	var reqBody SectionJSON
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
		return
	}

	if reqBody == (SectionJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
		return
	}

	sec := model.Section{
		ID:                 idInt,
		SectionNumber:      reqBody.SectionNumber,
		CurrentTemperature: reqBody.CurrentTemperature,
		MinimumTemperature: reqBody.MinimumTemperature,
		CurrentCapacity:    reqBody.CurrentCapacity,
		MinimumCapacity:    reqBody.MinimumCapacity,
		MaximumCapacity:    reqBody.MaximumCapacity,
		WarehouseID:        reqBody.WarehouseID,
		ProductTypeID:      reqBody.ProductTypeID,
	}

	s, err := h.sv.Update(idInt, sec)
	if err != nil {
		response.JSON(w, handleError(err), responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", s))
}

func (h *SectionController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		return
	}

	err = h.sv.Delete(idInt)
	if err != nil {
		response.JSON(w, handleError(err), responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("section removed successfully", nil))
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
