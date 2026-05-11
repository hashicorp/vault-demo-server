// Copyright IBM Corp. 2015, 2026

// +build appengine plan9

package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	return strings.Contains(err.Error(), "connection reset")
}
