// Copyright IBM Corp. 2015, 2026

// +build !go1.7

package pq

import "crypto/tls"

// Renegotiation is not supported by crypto/tls until Go 1.7.
func sslRenegotiation(*tls.Config) {}
