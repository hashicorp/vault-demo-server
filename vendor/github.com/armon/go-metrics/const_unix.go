// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build !windows

package metrics

import (
	"syscall"
)

const (
	// DefaultSignal is used with DefaultInmemSignal
	DefaultSignal = syscall.SIGUSR1
)
