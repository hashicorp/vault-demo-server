// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build !ui

package http

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func init() {
	uiBuiltIn = false
}

// assetFS is a stub for building Vault without a UI.
func assetFS() *assetfs.AssetFS {
	return nil
}
