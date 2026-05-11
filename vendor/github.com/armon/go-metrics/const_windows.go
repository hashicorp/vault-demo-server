// Copyright IBM Corp. 2015, 2026

// +build windows

package metrics

import (
	"syscall"
)

const (
	// DefaultSignal is used with DefaultInmemSignal
	// Windows has no SIGUSR1, use SIGBREAK
	DefaultSignal = syscall.Signal(21)
)
