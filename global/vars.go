package global

import "github.com/nsip/n3-messages/n3grpc"

var (
	// N3pub :
	N3pub *n3grpc.Publisher
)

type (
	SQType int
)

const (
	SIF       SQType = 0
	XAPI      SQType = 1
	META_SIF  SQType = 2
	META_XAPI SQType = 3
)
