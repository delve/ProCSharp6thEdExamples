package main

import (
	"errors"
	"fmt"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d directDeposit) report() {
	fmt.Println(d.lastName)
	fmt.Println(d.firstName)
	fmt.Println(d.bankName)
	fmt.Println(d.routingNumber)
	fmt.Println(d.accountNumber)
}

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	test := directDeposit{lastName: "", firstName: "Abe", bankName: "XYZ Inc", routingNumber: 17, accountNumber: 1809}
	if err := test.validateLastName(); err != nil {
		fmt.Println(err)
	}
	if err := test.validateRoutingNumber(); err != nil {
		fmt.Println(err)
	}
	test.report()
}
