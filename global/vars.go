package global

import "github.com/nsip/n3-messages/n3grpc"

var (
	// N3clt :
	N3clt *n3grpc.Client
)

type (
	SQDType int
)

const (
	SIF       SQDType = 0
	XAPI      SQDType = 1
	META_SIF  SQDType = 2
	META_XAPI SQDType = 3
)
