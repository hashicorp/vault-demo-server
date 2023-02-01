// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build windows

package pq

// sslKeyPermissions checks the permissions on user-supplied ssl key files.
// The key file should have very little access.
//
// libpq does not check key file permissions on Windows.
func sslKeyPermissions(string) error { return nil }
