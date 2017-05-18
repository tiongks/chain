// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearning

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Definition of the public APIs exposed by Amazon Machine Learning
type MachineLearning struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new MachineLearning client.
func New(config *aws.Config) *MachineLearning {
	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "machinelearning",
		APIVersion:   "2014-12-12",
		JSONVersion:  "1.1",
		TargetPrefix: "AmazonML_20141212",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &MachineLearning{service}
}

// newRequest creates a new request for a MachineLearning operation and runs any
// custom request initialization.
func (c *MachineLearning) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}