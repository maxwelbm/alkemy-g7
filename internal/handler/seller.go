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

func CreateHandlerSellers(service interfaces.ISellerService, log logger.Logger) *SellersController {
	return &SellersController{Service: service, log: log}
}

type SellersController struct {
	Service interfaces.ISellerService
	log     logger.Logger
}

// GetAllSellers retrieves all sellers.
// @Summary Retrieve all sellers
// @Description Fetch all registered sellers from the database
// @Tags Seller
// @Produce json
// @Success 200 {object} model.SellerResponseSwagger
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to list sellers"
// @Router /sellers [get]
func (hd *SellersController) GetAllSellers(w http.ResponseWriter, r *http.Request) {
	hd.log.Log("SellersHandler", "INFO", "Get all sellers initializing")

	sellers, err := hd.Service.GetAll()
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	data := make([]model.Seller, 0)
	for _, value := range sellers {
		data = append(data, model.Seller{
			ID:          value.ID,
			CID:         value.CID,
			CompanyName: value.CompanyName,
			Address:     value.Address,
			Telephone:   value.Telephone,
			Locality:    value.Locality,
		})
	}

	hd.log.Log("SellersHandler", "INFO", fmt.Sprintf("Retrieved sellers successfully: %+v", sellers))
	hd.log.Log("SellersHandler", "INFO", "Get all sellers completed")

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", data))
}

// GetByID retrieves a seller by their ID.
// @Summary Retrieve seller
// @Description This endpoint fetches the details of a specific seller based on the provided seller ID.
// @Tags Seller
// @Produce json
// @Param id path int true "Seller ID"
// @Success 200 {object} model.SellerResponseSwagger{data=model.Seller}
// @Failure 400 {object} model.ErrorResponseSwagger "missing 'id' parameter in the request"
// @Failure 404 {object} model.ErrorResponseSwagger "seller not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for seller"
// @Router /sellers/{id} [get]
func (hd *SellersController) GetByID(w http.ResponseWriter, r *http.Request) {
	hd.log.Log("SellersHandler", "INFO", "Get seller by ID initializing")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id == 0 || err != nil {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	seller, err := hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	hd.log.Log("SellersHandler", "INFO", fmt.Sprintf("Retrieved seller successfully: %+v", seller))
	hd.log.Log("SellersHandler", "INFO", "Get seller by ID completed")

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

// CreateSellers creates a new seller.
// @Summary Create a new seller
// @Description This endpoint allows for creating a new seller.
// @Tags Seller
// @Produce json
// @Param seller body model.Seller true "Seller information"
// @Success 201 {object} model.SellerResponseSwagger{data=model.Seller}
// @Failure 400 {object} model.ErrorResponseSwagger "Unprocessable Entity"
// @Failure 404 {object} model.ErrorResponseSwagger "Locality not found"
// @Failure 409 {object} model.ErrorResponseSwagger "CID number already exists"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create seller"
// @Router /sellers [post]
func (hd *SellersController) CreateSellers(w http.ResponseWriter, r *http.Request) {
	hd.log.Log("SellersHandler", "INFO", "Create sellers initializing")

	var seller model.Seller
	if err := request.JSON(r, &seller); err != nil {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		response.JSON(w, er.ErrInvalidSellerJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidSellerJSONFormat.Error(), nil))
		return
	}

	createdseller, err := hd.Service.CreateSeller(&seller)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	hd.log.Log("SellersHandler", "INFO", fmt.Sprintf("Created seller successfully: %+v", createdseller))
	hd.log.Log("SellersHandler", "INFO", "Create sellers completed")

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", createdseller))
}

// UpdateSellers updates an existing seller.
// @Summary Update an existing seller
// @Description This endpoint allows for updating the details of a specific seller identified by the provided ID.
// @Tags Seller
// @Produce json
// @Param id path int true "Seller ID"
// @Param seller body model.Seller true "Seller information"
// @Success 200 {object} model.SellerResponseSwagger{data=model.Seller} "Seller successfully updated"
// @Failure 422 {object} model.ErrorResponseSwagger "Unprocessable Entity"
// @Failure 404 {object} model.ErrorResponseSwagger "Seller not found"
// @Failure 404 {object} model.ErrorResponseSwagger "Locality not found"
// @Failure 409 {object} model.ErrorResponseSwagger "CID number already exists"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to update seller"
// @Router /sellers/{id} [patch]
func (hd *SellersController) UpdateSellers(w http.ResponseWriter, r *http.Request) {
	hd.log.Log("SellersHandler", "INFO", "Update sellers initializing")

	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)

	if id == 0 || err != nil {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	_, err = hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	var s model.Seller
	if err := request.JSON(r, &s); err != nil {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		response.JSON(w, er.ErrInvalidSellerJSONFormat.Code, responses.CreateResponseBody(er.ErrInvalidSellerJSONFormat.Error(), nil))
		return
	}

	seller, err := hd.Service.UpdateSeller(id, &s)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	hd.log.Log("SellersHandler", "INFO", fmt.Sprintf("Updated seller successfully: %+v", seller))
	hd.log.Log("SellersHandler", "INFO", "Update sellers completed")

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", seller))
}

// DeleteSellers deletes a seller by their ID.
// @Summary Delete a seller by ID
// @Description This endpoint allows for deleting a seller based on the provided seller ID.
// @Tags Seller
// @Produce json
// @Param id path int true "Seller ID"
// @Success 204 {object} nil "Seller successfully deleted"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Seller not found"
// @Failure 409 {object} model.ErrorResponseSwagger "Seller cannot be deleted due to existing dependencies"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to delete seller"
// @Router /sellers/{id} [delete]
func (hd *SellersController) DeleteSellers(w http.ResponseWriter, r *http.Request) {
	hd.log.Log("SellersHandler", "INFO", "Delete sellers initializing")

	idSearch := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSearch)

	if id == 0 || err != nil {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())
		err := er.ErrMissingSellerID
		response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

		return
	}

	_, err = hd.Service.GetByID(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	err = hd.Service.DeleteSeller(id)
	if ok := hd.handlerError(err, w); ok {
		hd.log.Log("SellersHandler", "ERROR", "Error: "+err.Error())

		return
	}

	hd.log.Log("SellersHandler", "INFO", fmt.Sprintf("Removed seller successfully with id: %+d", id))
	hd.log.Log("SellersHandler", "INFO", "Delete sellers completed")

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("", nil))
}

func (hd *SellersController) handlerError(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err, ok := err.(*er.SellerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return true
		}

		if err, ok := err.(*er.LocalityError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return true
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unmapped seller handler error", nil))

		return true
	}

	return false
}
