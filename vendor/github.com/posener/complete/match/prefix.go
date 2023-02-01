// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package match

import "strings"

// Prefix is a simple Matcher, if the word is it's prefix, there is a match
// Match returns true if a has the prefix as prefix
func Prefix(long, prefix string) bool {
	return strings.HasPrefix(long, prefix)
}
