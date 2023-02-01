// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//+build sasl

package mgo

import (
	"gopkg.in/mgo.v2/internal/sasl"
)

func saslNew(cred Credential, host string) (saslStepper, error) {
	return sasl.New(cred.Username, cred.Password, cred.Mechanism, cred.Service, host)
}
