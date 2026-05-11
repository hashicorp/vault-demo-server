// Copyright IBM Corp. 2015, 2026

package pki

import "strings"

func normalizeSerial(serial string) string {
	return strings.Replace(strings.ToLower(serial), ":", "-", -1)
}
