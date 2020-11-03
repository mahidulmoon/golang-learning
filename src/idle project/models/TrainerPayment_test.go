package models

import (
	"fmt"
	"testing"
)

func TestAddTrainerPayment(t *testing.T) {
	f := TrainerPayment{

		Trainer_Id :   3,
		Bkash :         "bkash check ",
		BankInfo : "dbbl",
		Total :         2500.00 ,
		Due :         100.00,
		Percentage :         34.50,
		Comments :       "this is a comments to check ",
	}
	err := f.Add()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println("test passed paymentinfo added.")

}

func TestGetTrainerPayment(t *testing.T) {
	f, err := GetAllTrainerPayments()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(f, err)

}

func TestGetTrainerPaymentbyTID(t *testing.T) {
	f, err := GetTrainerPaymentByTID(3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(f, err)

}
