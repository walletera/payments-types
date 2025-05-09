package events

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/google/uuid"
    "github.com/walletera/eventskit/events"
    "github.com/walletera/payments-types/pkg/wogen"
    "github.com/walletera/payments-types/pkg/wuuid"
    api "github.com/walletera/payments-types/privateapi"
    "github.com/walletera/werrors"
)

var _ events.Event[Handler] = PaymentCreated{}

const (
    PaymentCreatedType = "PaymentCreated"
)

type PaymentCreated struct {
    Id                 uuid.UUID      `json:"id"`
    EventType          string         `json:"type"`
    EventCorrelationID string         `json:"correlationId"`
    SerializableData   json.Marshaler `json:"data"`

    Data api.Payment `json:"-"`
}

func NewPaymentCreated(correlationId string, data api.Payment) PaymentCreated {
    return PaymentCreated{
        Id:                 wuuid.NewUUID(),
        EventType:          PaymentCreatedType,
        EventCorrelationID: correlationId,
        Data:               data,
        SerializableData:   wogen.NewSerializationWrapper(&data),
    }
}

func (w PaymentCreated) Accept(ctx context.Context, handler Handler) werrors.WError {
    return handler.HandlePaymentCreated(ctx, w)
}

func (w PaymentCreated) ID() string {
    return fmt.Sprintf("%s-%s", w.Type(), w.Id)
}

func (w PaymentCreated) Type() string {
    return PaymentCreatedType
}

func (w PaymentCreated) CorrelationID() string {
    return w.EventCorrelationID
}

func (w PaymentCreated) DataContentType() string {
    return "application/json"
}

func (w PaymentCreated) Serialize() ([]byte, error) {
    return json.Marshal(w)
}
