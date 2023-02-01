// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

// Constraint is used to serialize a job placement constraint.
type Constraint struct {
	LTarget string
	RTarget string
	Operand string
}

// NewConstraint generates a new job placement constraint.
func NewConstraint(left, operand, right string) *Constraint {
	return &Constraint{
		LTarget: left,
		RTarget: right,
		Operand: operand,
	}
}
