// Code generated by thriftrw-plugin-yarpc
// @generated

package storeserver

import (
	"context"
	"go.uber.org/thriftrw/envelope"
	"go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcerror"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/atomic"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/atomic/readonlystoreserver"
)

// Interface is the server-side interface for the Store service.
type Interface interface {
	readonlystoreserver.Interface

	CompareAndSwap(
		ctx context.Context,
		Request *atomic.CompareAndSwap,
	) error

	Increment(
		ctx context.Context,
		Key *string,
		Value *int64,
	) error
}

// New prepares an implementation of the Store service for
// registration.
//
// 	handler := StoreHandler{}
// 	dispatcher.Register(storeserver.New(handler))
func New(impl Interface, opts ...yarpcthrift.RegisterOption) []yarpc.EncodingProcedure {
	h := handler{impl}
	service := yarpcthrift.Service{
		Name: "Store",
		Methods: []yarpcthrift.Method{

			yarpcthrift.Method{
				Name:         "compareAndSwap",
				Handler:      yarpcthrift.EncodingHandler(h.CompareAndSwap),
				Signature:    "CompareAndSwap(Request *atomic.CompareAndSwap)",
				ThriftModule: atomic.ThriftModule,
			},

			yarpcthrift.Method{
				Name:         "increment",
				Handler:      yarpcthrift.EncodingHandler(h.Increment),
				Signature:    "Increment(Key *string, Value *int64)",
				ThriftModule: atomic.ThriftModule,
			},
		},
	}

	procedures := make([]yarpc.EncodingProcedure, 0, 3)
	procedures = append(procedures, readonlystoreserver.New(impl, opts...)...)
	procedures = append(procedures, yarpcthrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) CompareAndSwap(ctx context.Context, body wire.Value) (envelope.Enveloper, error) {
	var args atomic.Store_CompareAndSwap_Args
	if err := args.FromWire(body); err != nil {
		return nil, err
	}

	appErr := h.impl.CompareAndSwap(ctx, args.Request)

	result, err := atomic.Store_CompareAndSwap_Helper.WrapResponse(appErr)
	if err != nil {
		return nil, err
	}

	if appErr != nil {
		return nil, yarpcerror.New(
			yarpcerror.CodeUnknown,
			appErr.Error(),
			yarpcerror.WithDetails(result),
		)
	}
	return result, nil
}

func (h handler) Increment(ctx context.Context, body wire.Value) (envelope.Enveloper, error) {
	var args atomic.Store_Increment_Args
	if err := args.FromWire(body); err != nil {
		return nil, err
	}

	appErr := h.impl.Increment(ctx, args.Key, args.Value)

	result, err := atomic.Store_Increment_Helper.WrapResponse(appErr)
	if err != nil {
		return nil, err
	}

	if appErr != nil {
		return nil, yarpcerror.New(
			yarpcerror.CodeUnknown,
			appErr.Error(),
			yarpcerror.WithDetails(result),
		)
	}
	return result, nil
}
