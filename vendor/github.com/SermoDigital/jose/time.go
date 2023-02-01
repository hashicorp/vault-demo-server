// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jose

import "time"

// Now returns the current time in UTC.
func Now() time.Time { return time.Now().UTC() }
