package main

import (
	"errors"
	"log"
)

var (
	InvalidLastName      = errors.New("invalid last name")
	InvalidRoutingNumber = errors.New("invalid routing number")
)

type DirectDeposite struct {
	firstName, lastName          string
	routingNumber, accountNumber int
}

func makeDirectDeposite(dd struct {
	firstName, lastName          string
	routingNumber, accountNumber int
}) *DirectDeposite {
	return &DirectDeposite{dd.firstName, dd.lastName, dd.routingNumber, dd.accountNumber}
}

func (dd *DirectDeposite) validateRoutingNumber() error {
	if dd.routingNumber < 100 {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		panic(InvalidRoutingNumber)
	}
	return nil
}

func (dd *DirectDeposite) validateLastName() error {
	if dd.lastName == "" {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		panic(InvalidLastName)
	}
	return nil
}
