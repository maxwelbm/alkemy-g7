package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type BuyerHandler struct {
	Svc interfaces.IBuyerservice
}

func NewBuyerHandler(svc interfaces.IBuyerservice) *BuyerHandler {
	return &BuyerHandler{svc}
}

// HandlerGetAllBuyers retrieves all buyers.
// @Summary Retrieve all buyers
// @Description Fetch all registered buyers from the database
// @Tags Buyes
// @Produce json
// @Success 200 {object} model.BuyerResponseSwagger
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to list Buyers"
// @Router /buyers [get]
func (bh *BuyerHandler) HandlerGetAllBuyers(w http.ResponseWriter, r *http.Request) {
	buyers, err := bh.Svc.GetAllBuyer()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to list Buyers", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyers))
}

// HandlerGetBuyerByID retrieves a buyer by their ID.
// @Summary Retrieve buyer
// @Description This endpoint fetches the details of a specific buyer based on the provided buyer ID. It returns the buyer's information, including their name and any other relevant details. If the buyer ID does not exist, it returns a 404 Not Found error with an appropriate message.
// @Tags Buyer
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} model.BuyerResponseSwagger{data=model.Buyer}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Buyer Not Found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for buyer"
// @Router /buyers/{id} [get]
func (bh *BuyerHandler) HandlerGetBuyerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/Buyers/"):]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	buyer, err := bh.Svc.GetBuyerByID(id)

	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to search for buyer", nil))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyer))
}

// HandlerDeleteBuyerByID deletes a buyer by their ID.
// @Summary Delete a buyer by ID
// @Description This endpoint allows for deleting a buyer based on the provided buyer ID. It checks for the existence of the buyer and any dependencies that might prevent deletion.
// @Tags Buyer
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 204 {object} nil "Buyer successfully deleted"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Buyer not found"
// @Failure 409 {object} model.ErrorResponseSwagger "Buyer cannot be deleted due to existing dependencies"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to delete buyer"
// @Router /buyers/{id} [delete]
func (bh *BuyerHandler) HandlerDeleteBuyerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/Buyers/"):]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	err = bh.Svc.DeleteBuyerByID(id)

	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to delete buyer", nil))

		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// HandlerCreateBuyer creates a new buyer.
// @Summary Create a new buyer
// @Description This endpoint allows for creating a new buyer. It validates the input and checks for unique constraints on the card number.
// @Description 422 responses may include:
// @Description - JSON syntax error (malformed JSON).
// @Description - Mandatory fields not filled in.
// @Tags Buyer
// @Produce json
// @Param buyer body model.Buyer true "Buyer information"
// @Success 201 {object} model.BuyerResponseSwagger{data=model.Buyer}
// @Failure 400 {object} model.ErrorResponseSwagger "Unprocessable Entity"
// @Failure 409 {object} model.ErrorResponseSwagger "Card number already exists"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create buyer"
// @Router /buyers [post]
func (bh *BuyerHandler) HandlerCreateBuyer(w http.ResponseWriter, r *http.Request) {
	var reqBody model.Buyer

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))
		return
	}

	err = reqBody.ValidateEmptyFields(false)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	buyer, err := bh.Svc.CreateBuyer(reqBody)

	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create buyer", nil))

		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", buyer))
}

// HandlerUpdateBuyer updates an existing buyer.
// @Summary Update an existing buyer
// @Description This endpoint allows for updating the details of a specific buyer identified by the provided ID. It validates the input and checks for unique constraints on the card number.
// @Description This endpoint performs the following actions:
// @Description 1. Validates the provided ID and ensures it corresponds to an existing buyer.
// @Description 2. Validates the input JSON for correct structure and required fields.
// @Description 3. Checks for unique constraints, such as unique card numbers.
// @Description Responses for errors may include:
// @Description - **422**: Unprocessable Entity, responses may include:
// @Description   - JSON syntax error (malformed JSON).
// @Description   - Mandatory fields not filled in.
// @Description
// @Description - **404**: Buyer not found, indicating the specified buyer does not exist.
// @Description - **409**: Card number already exists, indicating a unique constraint violation.
// @Description  - **500**: Internal server error for unexpected issues.
// @Tags Buyer
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body model.Buyer true "Buyer information"
// @Success 200 {object} model.BuyerResponseSwagger{data=model.Buyer} "Buyer successfully updated"
// @Example 200 { "data": {"id": 1, "name": "Updated Buyer", "card_number": "1234-5678-9012-3456"} }
// @Failure 422 {object} model.ErrorResponseSwagger "Unprocessable Entity"
// @Failure 404 {object} model.ErrorResponseSwagger "Buyer not found"
// @Failure 409 {object} model.ErrorResponseSwagger "Card number already exists"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to update buyer"
// @Router /buyers/{id} [patch]
func (bh *BuyerHandler) HandlerUpdateBuyer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/Buyers/"):]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	var reqBody model.Buyer

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("JSON syntax error. Please verify your input.", nil))
		return
	}

	err = reqBody.ValidateEmptyFields(true)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	buyer, err := bh.Svc.UpdateBuyer(id, reqBody)

	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to update buyer", nil))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", buyer))
}

// HandlerCountPurchaseOrderBuyer counts the purchase orders for a buyer.
// @Summary Count purchase orders for a buyer
// @Description This endpoint retrieves the count of purchase orders for a buyer. If an ID is not provided, it returns the total count of all purchase orders.
// @Tags Buyer
// @Produce json
// @Param id query int false "Buyer ID"
// @Success 200 {object} model.BuyerResponseSwagger{data=model.BuyerPurchaseOrder}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to count buyer purchase orders"
// @Router /buyers/reportPurchaseOrders [get]
func (bh *BuyerHandler) HandlerCountPurchaseOrderBuyer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		count, err := bh.Svc.CountPurchaseOrderBuyer()
		if err != nil {
			if err, ok := err.(*customerror.BuyerError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to count buyer Purchase orders", nil))

			return
		}

		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid ID", nil))
		return
	}

	count, err := bh.Svc.CountPurchaseOrderByBuyerID(id)
	if err != nil {
		if err, ok := err.(*customerror.BuyerError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to count buyer Purchase orders", nil))

		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", count))
}
