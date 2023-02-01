// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build !go1.6

package s3

import "github.com/aws/aws-sdk-go/aws/request"

func platformRequestHandlers(r *request.Request) {
}
