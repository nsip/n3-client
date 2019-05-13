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
	SIF SQDType = iota
	XAPI
	META_SIF
	META_XAPI
)
