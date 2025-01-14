package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type CarrierHandler struct {
	srv svc.ICarrierService
}

func NewCarrierHandler(srv svc.ICarrierService) *CarrierHandler {
	return &CarrierHandler{srv: srv}
}

func (h *CarrierHandler) PostCarriers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody model.Carries

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err := reqBody.ValidateEmptyFields(false)

		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		carrier, err := h.srv.PostCarrier(reqBody)

		if err != nil {
			if err, ok := err.(*custom_error.CarrierError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			if err, ok := err.(*custom_error.LocalityError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to post carrier", nil))
			return
		}

		response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", carrier))
	}
}
