package handler

import (
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

func CreateHandlerLocality(service interfaces.ILocalityService, log logger.Logger) *LocalitiesController {
	return &LocalitiesController{Service: service, log: log}
}

type LocalitiesController struct {
	Service interfaces.ILocalityService
	log     logger.Logger
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
	hd.log.Log("LocalitiesHandler", "INFO", "Get locality by ID initializing")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id == 0 || err != nil {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))
		err := er.ErrMissingLocalityID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	locality, err := hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	hd.log.Log("LocalitiesHandler", "INFO", fmt.Sprintf("Retrieved locality successfully: %+v", locality))
	hd.log.Log("LocalitiesHandler", "INFO", "Get locality by ID completed")

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
	hd.log.Log("LocalitiesHandler", "INFO", "Create locality initializing")

	var locality model.Locality
	if err := request.JSON(r, &locality); err != nil {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))
		response.JSON(w, er.ErrInvalidLocalityJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidLocalityJSONFormat.Error(), nil))

		return
	}

	createdLocality, err := hd.Service.CreateLocality(&locality)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	hd.log.Log("LocalitiesHandler", "INFO", fmt.Sprintf("Created locality successfully: %+v", createdLocality))
	hd.log.Log("LocalitiesHandler", "INFO", "Create locality completed")

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
	hd.log.Log("LocalitiesHandler", "INFO", "Get report Sellers initializing")

	id := 0

	if len(r.URL.Query()) > 0 {
		param := r.URL.Query().Get("id")
		if param == "" {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", er.ErrMissingLocalityID))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrMissingLocalityID.Error(), nil))

			return
		}

		idParam, err := strconv.Atoi(param)
		if idParam == 0 {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", er.ErrInvalidLocalityPathParam))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrInvalidLocalityPathParam.Error(), nil))

			return
		}

		if ok := hd.handlerError(err, w); ok {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		id = idParam
	}

	result, err := hd.Service.GetSellers(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	hd.log.Log("LocalitiesHandler", "INFO", fmt.Sprintf("Get report sellers successfully: %+v", result))
	hd.log.Log("LocalitiesHandler", "INFO", "Get report Sellers completed")

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
	hd.log.Log("LocalitiesHandler", "INFO", "Get report Carriers initializing")

	id := 0

	if len(r.URL.Query()) > 0 {
		param := r.URL.Query().Get("id")
		if param == "" {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", er.ErrMissingLocalityID))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrMissingLocalityID.Error(), nil))

			return
		}

		idParam, err := strconv.Atoi(param)
		if idParam == 0 {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", er.ErrInvalidLocalityPathParam))
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody(er.ErrInvalidLocalityPathParam.Error(), nil))

			return
		}

		if ok := hd.handlerError(err, w); ok {
			hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		id = idParam

	}

	result, err := hd.Service.GetCarriers(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("LocalitiesHandler", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	hd.log.Log("LocalitiesHandler", "INFO", fmt.Sprintf("Get report carriers successfully: %+v", result))
	hd.log.Log("LocalitiesHandler", "INFO", "Get report Carriers completed")

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
