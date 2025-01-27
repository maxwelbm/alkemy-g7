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

// GetByID retrieves a locality by their ID.
// @Summary Retrieve locality
// @Description This endpoint fetches the details of a specific locality based on the provided locality ID.
// @Tags Locality
// @Produce json
// @Param id path int true "Locality ID"
// @Success 200 {object} model.LocalityResponseSwagger{data=model.Locality}
// @Failure 400 {object} model.ErrorResponseSwagger "missing 'id' parameter in the request"
// @Failure 404 {object} model.ErrorResponseSwagger "locality not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for locality"
// @Router /localities/{id} [get]
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

// CreateLocality creates a new locality.
// @Summary Create a new locality
// @Description This endpoint allows for creating a new locality.
// @Tags Locality
// @Produce json
// @Param locality body model.Locality true "Locality information"
// @Success 201 {object} model.LocalityResponseSwagger{data=model.Locality}
// @Failure 400 {object} model.ErrorResponseSwagger "Unprocessable Entity"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create locality"
// @Router /localities [post]
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

// GetSellers retrieves a count of sellers from locality by their ID.
// @Summary Retrieve locality and count sellers
// @Description This endpoint fetches the details of a specific locality abount sellers count based on the provided locality ID.
// @Tags Locality
// @Produce json
// @Param id path int false "Locality ID"
// @Success 200 {object} model.LocalitySellersResponseSwagger{data=model.LocalitiesJSONSellers}
// @Failure 400 {object} model.ErrorResponseSwagger "missing 'id' parameter in the request"
// @Failure 404 {object} model.ErrorResponseSwagger "locality not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for locality"
// @Router /localities/reportSellers [get]
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

// GetCarriers retrieves a count of carriers from locality by their ID.
// @Summary Retrieve locality and count carriers
// @Description This endpoint fetches the details of a specific locality abount carriers count based on the provided locality ID.
// @Tags Locality
// @Produce json
// @Param id path int false "Locality ID"
// @Success 200 {object} model.LocalityCarriersResponseSwagger{data=model.LocalitiesJSONCarriers}
// @Failure 400 {object} model.ErrorResponseSwagger "missing 'id' parameter in the request"
// @Failure 404 {object} model.ErrorResponseSwagger "locality not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for locality"
// @Router /localities/reportCarriers [get]
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
