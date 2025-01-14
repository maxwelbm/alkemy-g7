package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

func CreateHandlerLocality(service interfaces.ILocalityService) *LocalitiesController {
	return &LocalitiesController{service: service}
}

type LocalitiesController struct {
	service interfaces.ILocalityService
}

func (hd *LocalitiesController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorMissingLocalityID.Error(), nil))
		return
	}

	locality, err := hd.service.GetById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", locality))
}

func (hd *LocalitiesController) CreateLocality(w http.ResponseWriter, r *http.Request) {
	var locality model.Locality
	if err := request.JSON(r, &locality); err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorInvalidLocalityJSONFormat.Error(), nil))
		return
	}

	createdLocality, err := hd.service.CreateLocality(&locality)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdLocality))
}

func (hd *LocalitiesController) GetSellers(w http.ResponseWriter, r *http.Request) {
	var id int = 0

	if len(r.URL.Query()) > 0 {
		existID := r.URL.Query().Has("id")
		if !existID {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorMissingSellerID.Error(), nil))
			return
		}
		param := r.URL.Query().Get("id")
		if param != "" {
			idParam, err := strconv.Atoi(param)
			id = idParam
			if err != nil {
				response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorInvalidLocalityPathParam.Error(), nil))
				return
			}
		}
	}

	result, err := hd.service.GetSellers(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", result))
}

func (hd *LocalitiesController) GetCarriers(w http.ResponseWriter, r *http.Request) {
	var id int = 0

	if len(r.URL.Query()) > 0 {
		existID := r.URL.Query().Has("id")
		if !existID {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorMissingSellerID.Error(), nil))
			return
		}
		param := r.URL.Query().Get("id")
		if param != "" {
			idParam, err := strconv.Atoi(param)
			id = idParam
			if err != nil {
				response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(model.ErrorInvalidLocalityPathParam.Error(), nil))
				return
			}
		}
	}

	result, err := hd.service.GetCarriers(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", result))
}
