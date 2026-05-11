// Copyright IBM Corp. 2015, 2026

package plugin

import (
	"math"

	"google.golang.org/grpc"
)

var largeMsgGRPCCallOpts []grpc.CallOption = []grpc.CallOption{
	grpc.MaxCallSendMsgSize(math.MaxInt32),
	grpc.MaxCallRecvMsgSize(math.MaxInt32),
}
