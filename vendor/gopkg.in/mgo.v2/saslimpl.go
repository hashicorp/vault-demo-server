// Copyright IBM Corp. 2015, 2026

//+build sasl

package mgo

import (
	"gopkg.in/mgo.v2/internal/sasl"
)

func saslNew(cred Credential, host string) (saslStepper, error) {
	return sasl.New(cred.Username, cred.Password, cred.Mechanism, cred.Service, host)
}
