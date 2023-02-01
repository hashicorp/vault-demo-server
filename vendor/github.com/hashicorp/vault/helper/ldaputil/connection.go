// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ldaputil

import (
	"crypto/tls"

	"github.com/go-ldap/ldap"
)

// Connection provides the functionality of an LDAP connection,
// but through an interface.
type Connection interface {
	Bind(username, password string) error
	Close()
	Modify(modifyRequest *ldap.ModifyRequest) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	StartTLS(config *tls.Config) error
	UnauthenticatedBind(username string) error
}
