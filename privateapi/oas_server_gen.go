// Code generated by ogen, DO NOT EDIT.

package privateapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetPayment implements get-payment operation.
	//
	// Retrieves an individual payment.
	//
	// GET /payments/{paymentId}
	GetPayment(ctx context.Context, params GetPaymentParams) (GetPaymentRes, error)
	// PatchPayment implements patch-payment operation.
	//
	// Updates an outbound payment.
	//
	// PATCH /payments/{paymentId}
	PatchPayment(ctx context.Context, req *PaymentUpdate, params PatchPaymentParams) (PatchPaymentRes, error)
	// PostPayment implements post-payment operation.
	//
	// Creates a payment.
	//
	// POST /payments
	PostPayment(ctx context.Context, req *PostPaymentReq, params PostPaymentParams) (PostPaymentRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
