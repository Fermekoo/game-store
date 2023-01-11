package payment

import (
	"fmt"

	"github.com/Fermekoo/game-store/utils"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var mdCore coreapi.Client

type MidtransPay struct {
}

func NewMidtrans(config utils.Config) *MidtransPay {
	mdCore.New(config.MidtransServerKey, midtrans.Sandbox)
	return &MidtransPay{}
}

func (m *MidtransPay) Charge(payloads *ChargeRequest) (*ChargeResponse, error) {
	var chargeReq *coreapi.ChargeReq
	bank := midtrans.Bank(payloads.Bank)

	if payloads.PaymentType == "gopay" {
		chargeReq = &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  payloads.OrderID,
				GrossAmt: payloads.Amount,
			},
			Gopay: &coreapi.GopayDetails{
				EnableCallback: true,
				CallbackUrl:    "someapps://callback",
			},
			CustomerDetails: &midtrans.CustomerDetails{
				FName: payloads.Name,
				Phone: payloads.Phone,
			},
		}
	} else {
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeBankTransfer,
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: bank,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  payloads.OrderID,
				GrossAmt: payloads.Amount,
			},
		}
	}

	charge, err := mdCore.ChargeTransaction(chargeReq)

	if err != nil {
		return nil, fmt.Errorf(err.Message)
	}
	var payment_code string

	if charge.PaymentType == string(coreapi.PaymentTypeBankTransfer) {
		if bank == "permata" {
			payment_code = charge.PermataVaNumber
		} else if bank == "mandiri" {
			payment_code = fmt.Sprintf("%s%s", charge.BillerCode, charge.BillKey)
		} else {
			payment_code = charge.VaNumbers[0].VANumber
		}
	}

	result := &ChargeResponse{
		TransactionID: charge.TransactionID,
		OrderID:       charge.OrderID,
		PaymentType:   charge.PaymentType,
		PaymentCode:   payment_code,
	}

	return result, nil
}
