package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostInboundOrder(t *testing.T) {
	createRequest := func(body string) *http.Request {
		req := httptest.NewRequest("POST", "/api/v1/inbound-orders", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}

	srv := mocks.NewMockIInboundOrderService(t)
	handler := NewInboundHandler(srv, mocks.MockLog{})

	newInboundOrder := InboundOrderJSON{
		ID:             1,
		OrderDate:      "2023-10-01",
		OrderNumber:    "ORD123",
		EmployeeID:     1,
		ProductBatchID: 1,
		WareHouseID:    1,
	}

	inboundOrderJSON, err := json.Marshal(newInboundOrder)
	if err != nil {
		t.Fatalf("failed to marshal new inbound order: %v", err)
	}

	t.Run("should return 201 created and the new inbound order created", func(t *testing.T) {
		mockInboundOrder := model.InboundOrder{
			ID:             1,
			OrderDate:      time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			OrderNumber:    "ORD123",
			EmployeeID:     1,
			ProductBatchID: 1,
			WareHouseID:    1,
		}
		srv.On("Post", mockInboundOrder).Return(mockInboundOrder, nil).Once()
		req := createRequest(string(inboundOrderJSON))
		res := httptest.NewRecorder()

		handler.PostInboundOrder(res, req)

		expected := `{"message":"success","data":{"id":1,"order_date":"2023-10-01","order_number":"ORD123","employee_id":1,"product_batch_id":1,"warehouse_id":1}}`

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 400 bad request when invalid request body", func(t *testing.T) {
		reqBody := `something`
		req := createRequest(string(reqBody))
		res := httptest.NewRecorder()

		handler.PostInboundOrder(res, req)

		expected := `{"message":"error parsing the request body"}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 422 unprocessable entity when the input is missing fields", func(t *testing.T) {
		srv.On("Post", mock.Anything).Return(model.InboundOrder{}, customerror.NewInboundOrderErr("invalid input", http.StatusUnprocessableEntity)).Once()
		newInboundOrder := `
		{
			"order_number": "ORD123"
		}`

		req := createRequest(newInboundOrder)
		res := httptest.NewRecorder()

		handler.PostInboundOrder(res, req)

		expected := `{"message":"invalid input"}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 409 conflict when order number already exists", func(t *testing.T) {
		srv.On("Post", mock.Anything).Return(model.InboundOrder{}, customerror.NewInboundOrderErr("duplicated order number", http.StatusConflict)).Once()

		req := createRequest(string(inboundOrderJSON))
		res := httptest.NewRecorder()

		handler.PostInboundOrder(res, req)

		expected := `{"message":"duplicated order number"}`

		assert.Equal(t, http.StatusConflict, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		srv.On("Post", mock.Anything).Return(model.InboundOrder{}, errors.New("unexpected error")).Once()

		req := createRequest(string(inboundOrderJSON))
		res := httptest.NewRecorder()

		handler.PostInboundOrder(res, req)
		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.JSONEq(t, res.Body.String(), expected)
	})
}
