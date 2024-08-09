package payments

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/walletera/message-processor/events"
    "github.com/walletera/payments-types/api"
)

type EventsDeserializer struct {
}

func NewEventsDeserializer() *EventsDeserializer {
    return &EventsDeserializer{}
}

func (d *EventsDeserializer) Deserialize(rawPayload []byte) (events.Event[EventsHandler], error) {
    var event events.EventEnvelope
    err := json.Unmarshal(rawPayload, &event)
    if err != nil {
        return nil, fmt.Errorf("error deserializing message with payload %s: %w", rawPayload, err)
    }
    switch event.Type {
    case "PaymentCreated":
        var payment api.Payment
        err := json.Unmarshal(event.Data, &payment)
        if err != nil {
            log.Printf("error deserializing PaymentCreated event data %s: %s", event.Data, err.Error())
        }
        return NewPaymentCreatedEvent(event.CorrelationID, payment), nil
    default:
        log.Printf("unexpected event type: %s", event.Type)
        return nil, nil
    }
}
