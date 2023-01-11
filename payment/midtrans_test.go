package payment

import (
	"fmt"
	"log"
	"testing"

	"github.com/Fermekoo/game-store/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCharge(t *testing.T) {
	config, err := utils.LoadConfig("./../")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	midtrans := NewMidtrans(config)

	request := &ChargeRequest{
		PaymentType: "bank_transfer",
		Bank:        "permata",
		Amount:      10000,
		Name:        "dandi",
		Phone:       "081219344136",
		OrderID:     uuid.New().String(),
	}

	charge, err := midtrans.Charge(request)
	fmt.Println(charge)
	require.NoError(t, err)
	require.NotEmpty(t, charge)
	require.Equal(t, request.OrderID, charge.OrderID)
}
