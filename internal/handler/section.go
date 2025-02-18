package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
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

func CreateHandlerSections(sv interfaces.ISectionService, log logger.Logger) *SectionController {
	return &SectionController{Sv: sv, log: log}
}

type SectionController struct {
	Sv  interfaces.ISectionService
	log logger.Logger
}

func (h *SectionController) GetAll(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing GetAll controller function")
	s, err := h.Sv.Get()

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to list sections", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

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

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", data))
	h.log.Log("SectionController", "INFO", "returning a slice with all sections in JSON format")
}

func (h *SectionController) GetByID(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing GetByID controller function")

	idStr := r.URL.Path[len("/api/v1/sections/"):]
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	s, err := h.Sv.GetByID(idInt)
	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to search for section", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", s))
	h.log.Log("SectionController", "INFO", "returning a section in JSON format")
}

func (h *SectionController) Post(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing Post controller function")

	var reqBody SectionJSON

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid request body", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	if reqBody == (SectionJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

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

	s, err := h.Sv.Post(&section)
	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to create section", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", s))
	h.log.Log("SectionController", "INFO", "post function executed successfully")
}

func (h *SectionController) Update(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing Update controller function")

	idStr := r.URL.Path[len("/api/v1/sections/"):]
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	var reqBody SectionJSON
	err = json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	if reqBody == (SectionJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

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

	s, err := h.Sv.Update(idInt, &sec)
	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to update section", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", s))
	h.log.Log("SectionController", "INFO", "updated a section successfully")
}

func (h *SectionController) Delete(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing Delete controller function")

	idStr := r.URL.Path[len("/api/v1/sections/"):]
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id param", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	err = h.Sv.Delete(idInt)
	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to delete section", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("section removed successfully", nil))
	h.log.Log("SectionController", "INFO", "delete a section successfully")
}

func (h *SectionController) CountProductBatchesSections(w http.ResponseWriter, r *http.Request) {
	h.log.Log("SectionController", "INFO", "initializing CountProductBatchesSections controller function")

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		count, err := h.Sv.CountProductBatchesSections()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to count section product batches", nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	count, err := h.Sv.CountProductBatchesBySectionID(id)
	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("error", nil))
		h.log.Log("SectionController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))
	h.log.Log("SectionController", "INFO", "CountProductBatchesSections executed successfully")
}
