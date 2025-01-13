package handler

import (
	"errors"
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
		if ok := errors.Is(err, model.ErrorIDAlreadyExist); ok {
			response.JSON(w, http.StatusConflict, responses.CreateResponseBody(err.Error(), nil))
			return
		} else {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}
	}

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("", createdLocality))

}