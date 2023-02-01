// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// +build appengine

package murmur

import "encoding/binary"

func getBlock(data []byte, n int) (int64, int64) {
	k1 := binary.LittleEndian.Int64(data[n*16:])
	k2 := binary.LittleEndian.Int64(data[(n*16)+8:])
	return k1, k2
}
