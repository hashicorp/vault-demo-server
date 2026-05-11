// Copyright IBM Corp. 2015, 2026

//+build !sasl

package mgo

import (
	"fmt"
)

func saslNew(cred Credential, host string) (saslStepper, error) {
	return nil, fmt.Errorf("SASL support not enabled during build (-tags sasl)")
}
