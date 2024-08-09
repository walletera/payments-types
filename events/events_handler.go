package payments

import (
    "context"

    "github.com/walletera/message-processor/errors"
)

type EventsHandler interface {
    HandlePaymentCreated(ctx context.Context, paymentCreatedEvent PaymentCreated) errors.ProcessingError
}
