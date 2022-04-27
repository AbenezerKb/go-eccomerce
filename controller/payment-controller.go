package controller

import (
	"net/http"

	"context"
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func (c usercontroller) Payment(ctx *gin.Context) {
	var payment entity.Charge
	contx := ctx.Request.Context()
	ctx.BindJSON(&payment)

	apiKey := "sk_test_51KnGMSJbf18WVmMyYCHBH2AY7B7uTFLCzF7NktVQZ4ovrtfUSVnn8F7IgNj59hWV6bZEs0kFyQj5TzoPzdNSIY1Z00EkN0uzaI"
	stripe.Key = apiKey
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(payment.ProductName),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail)})

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_SAVE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_SAVE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	service.SavePayment(contx, &payment)

}
