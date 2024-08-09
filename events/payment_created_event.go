package payments

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/google/uuid"
    "github.com/walletera/message-processor/errors"
    "github.com/walletera/message-processor/events"
    "github.com/walletera/payments-types/api"
    "github.com/walletera/payments-types/pkg/wuuid"
)

var _ events.Event[EventsHandler] = PaymentCreated{}

type PaymentCreated struct {
    Id            uuid.UUID   `json:"id"`
    CorrelationId string      `json:"correlationId"`
    Data          api.Payment `json:"data"`
}

func NewPaymentCreatedEvent(correlationId string, data api.Payment) *PaymentCreated {
    return &PaymentCreated{
        Id:            wuuid.NewUUID(),
        CorrelationId: correlationId,
        Data:          data,
    }
}

func (w PaymentCreated) Accept(ctx context.Context, visitor EventsHandler) errors.ProcessingError {
    return visitor.HandlePaymentCreated(ctx, w)
}

func (w PaymentCreated) ID() string {
    return fmt.Sprintf("%s-%s", w.Type(), w.Id)
}

func (w PaymentCreated) Type() string {
    return "PaymentCreated"
}

func (w PaymentCreated) CorrelationID() string {
    return w.CorrelationId
}

func (w PaymentCreated) DataContentType() string {
    return "application/json"
}

func (w PaymentCreated) Serialize() ([]byte, error) {
    return json.Marshal(w)
}
