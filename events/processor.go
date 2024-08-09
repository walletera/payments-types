package payments

import "github.com/walletera/message-processor/messages"

// NewProcessor returns a messages.Processor which is specific to payments events
func NewProcessor(messageConsumer messages.Consumer, eventsVisitor EventsHandler, opts ...messages.ProcessorOpt) *messages.Processor[EventsHandler] {
    return messages.NewProcessor[EventsHandler](
        messageConsumer,
        NewEventsDeserializer(),
        eventsVisitor,
        opts...,
    )
}
