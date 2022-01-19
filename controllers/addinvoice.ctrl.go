package controllers

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/getAlby/lndhub.go/db/models"
	"github.com/getAlby/lndhub.go/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

// AddInvoiceController : Add invoice controller struct
type AddInvoiceController struct {
	svc *lib.LndhubService
}

func NewAddInvoiceController(svc *lib.LndhubService) *AddInvoiceController {
	return &AddInvoiceController{svc: svc}
}

// AddInvoice : Add invoice Controller
func (controller *AddInvoiceController) AddInvoice(c echo.Context) error {
	userID := c.Get("UserID").(int64)
	type RequestBody struct {
		Amt             uint   `json:"amt" validate:"required"`
		Memo            string `json:"memo"`
		DescriptionHash string `json:"description_hash"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind json, amt field with positive value is required",
		})
	}

	if err := c.Validate(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "amt with positive value is required",
		})
	}

	invoice := models.Invoice{
		Type:               "",
		UserID:             userID,
		TransactionEntryID: 0,
		Amount:             body.Amt,
		Memo:               body.Memo,
		DescriptionHash:    body.DescriptionHash,
		PaymentRequest:     "",
		RHash:              "",
		State:              "",
	}

	// TODO: move this to a service layer and call a method
	_, err := controller.svc.DB.NewInsert().Model(&invoice).Exec(context.TODO())
	if err != nil {
		c.Logger().Errorf("error saving an invoice: %v", err)
		// TODO: better error handling, possibly panic and catch in an error handler
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var responseBody struct {
		RHash          string `json:"r_hash"`
		PaymentRequest string `json:"payment_request"`
		PayReq         string `json:"pay_req"`
	}

	responseBody.PayReq = makePreimageHex()

	return c.JSON(http.StatusOK, &responseBody)
}

const hexBytes = random.Hex

func makePreimageHex() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = hexBytes[rand.Intn(len(hexBytes))]
	}
	return string(b)
}
