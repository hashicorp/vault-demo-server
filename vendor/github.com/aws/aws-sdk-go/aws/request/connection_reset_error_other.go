// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build appengine plan9

package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	return strings.Contains(err.Error(), "connection reset")
}
