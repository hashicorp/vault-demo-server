// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pki

import "strings"

func normalizeSerial(serial string) string {
	return strings.Replace(strings.ToLower(serial), ":", "-", -1)
}
