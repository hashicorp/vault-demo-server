// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../cmd/radius-dict-gen/main.go -package rfc2865 -output generated.go /usr/share/freeradius/dictionary.rfc2865

package rfc2865
