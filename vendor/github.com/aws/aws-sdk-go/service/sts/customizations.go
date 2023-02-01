// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sts

import "github.com/aws/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opAssumeRoleWithSAML, opAssumeRoleWithWebIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
