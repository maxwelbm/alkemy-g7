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
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

func CreateHandlerLocality(service interfaces.ILocalityService) *LocalitiesController {
	return &LocalitiesController{service: service}
}

type LocalitiesController struct {
	service interfaces.ILocalityService
}

func (hd *LocalitiesController) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id == 0 || err != nil {
		err := er.ErrMissingLocalityID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	if ok := hd.handlerError(err, w); ok {
		return
	}

	locality, err := hd.service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", locality))
}

func (hd *LocalitiesController) CreateLocality(w http.ResponseWriter, r *http.Request) {
	var locality model.Locality
	if err := request.JSON(r, &locality); err != nil {
		response.JSON(w, er.ErrInvalidLocalityJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidLocalityJSONFormat.Error(), nil))
		return
	}

	createdLocality, err := hd.service.CreateLocality(&locality)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdLocality))
}

func (hd *LocalitiesController) GetSellers(w http.ResponseWriter, r *http.Request) {
	id := 0

	if len(r.URL.Query()) > 0 {
		param := r.URL.Query().Get("id")
		if param == "" {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrMissingLocalityID.Error(), nil))
			return
		}

		if param != "" {
			idParam, err := strconv.Atoi(param)
			if idParam == 0 {
				response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrInvalidLocalityPathParam.Error(), nil))
				return
			}

			if ok := hd.handlerError(err, w); ok {
				return
			}

			id = idParam
		}
	}

	result, err := hd.service.GetSellers(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", result))
}

func (hd *LocalitiesController) GetCarriers(w http.ResponseWriter, r *http.Request) {
	id := 0

	if len(r.URL.Query()) > 0 {
		param := r.URL.Query().Get("id")
		if param == "" {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrMissingLocalityID.Error(), nil))
			return
		}

		if param != "" {
			idParam, err := strconv.Atoi(param)
			if idParam == 0 {
				response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrInvalidLocalityPathParam.Error(), nil))
				return
			}

			if ok := hd.handlerError(err, w); ok {
				return
			}

			id = idParam
		}
	}

	result, err := hd.service.GetCarriers(id)
	if ok := hd.handlerError(err, w); ok {
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", result))
}

func (hd *LocalitiesController) handlerError(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err, ok := err.(*er.LocalityError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return true
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unmapped locality handler error", nil))

		return true
	}

	return false
}
