package events

import (
    "bytes"
    "context"
    "log/slog"
    "os"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

func TestDeserializer_Deserialize(t *testing.T) {
    logger := slog.New(slog.NewJSONHandler(new(bytes.Buffer), nil))
    deserializer := NewDeserializer(logger)

    t.Run("PaymentCreated with all fields completed", func(t *testing.T) {
        jsonPayload, err := os.ReadFile("testdata/payment_created_dinopay_full.json")
        require.NoError(t, err)

        event, err := deserializer.Deserialize(jsonPayload)
        require.NoError(t, err)

        ctx := context.TODO()

        mockHandler := NewMockHandler(t)
        mockHandler.EXPECT().HandlePaymentCreated(ctx, event).Return(nil).Run(
            func(ctx context.Context, paymentCreatedEvent PaymentCreated) {
                // Event fields
                assert.Equal(t, "3b3315ea-38c1-40a4-b7f9-149cc9807096", paymentCreatedEvent.Id.String())
                assert.Equal(t, "PaymentCreated", paymentCreatedEvent.EventType)
                assert.Equal(t, uint64(1), paymentCreatedEvent.EventAggregateVersion)
                assert.Equal(t, "12345678-1234-1234-1234-123456789012", paymentCreatedEvent.EventCorrelationId)
                assert.Equal(t, "2021-01-01T00:00:00Z", paymentCreatedEvent.EventCreatedAt.Format(time.RFC3339))

                // Payment fields
                assert.Equal(t, "3b3315ea-38c1-40a4-b7f9-149cc9807096", paymentCreatedEvent.Data.ID.String())
                assert.Equal(t, float64(100), paymentCreatedEvent.Data.Amount)
                assert.Equal(t, "ARS", string(paymentCreatedEvent.Data.Currency))
                assert.Equal(t, "dinopay", string(paymentCreatedEvent.Data.Gateway))
                assert.Equal(t, "inbound", string(paymentCreatedEvent.Data.Direction))
                assert.Equal(t, "997e76a7-c18d-4337-90ea-9a5cb0c5b54e", paymentCreatedEvent.Data.CustomerId.String())
                assert.Equal(t, "confirmed", string(paymentCreatedEvent.Data.Status))
                assert.Equal(t, "2021-01-01T00:00:00Z", paymentCreatedEvent.Data.CreatedAt.Format(time.RFC3339))
                assert.Equal(t, "2021-01-01T00:00:00Z", paymentCreatedEvent.Data.UpdatedAt.Format(time.RFC3339))

                // Debtor fields
                assert.True(t, paymentCreatedEvent.Data.Debtor.InstitutionName.IsSet(), "Debtor InstitutionName should be set")
                assert.Equal(t, "dinopay", paymentCreatedEvent.Data.Debtor.InstitutionName.Value)
                assert.True(t, paymentCreatedEvent.Data.Debtor.InstitutionId.IsSet(), "Debtor InstitutionId should be set")
                assert.Equal(t, "dinopay", paymentCreatedEvent.Data.Debtor.InstitutionId.Value)
                assert.Equal(t, "ARS", string(paymentCreatedEvent.Data.Debtor.Currency))

                // Check if Debtor AccountDetails is CvuAccountDetails
                debtorCvuDetails, ok := paymentCreatedEvent.Data.Debtor.AccountDetails.OneOf.GetCvuAccountDetails()
                assert.True(t, ok, "Debtor AccountDetails should be CvuAccountDetails")
                assert.Equal(t, "cvu", debtorCvuDetails.AccountType)
                assert.True(t, debtorCvuDetails.Cuit.IsSet(), "Debtor Cuit should be set")
                assert.Equal(t, "23679876453", debtorCvuDetails.Cuit.Value)

                // Check if Debtor RoutingInfo is CvuCvuRoutingInfo
                debtorCvuRoutingInfo, ok := debtorCvuDetails.RoutingInfo.OneOf.GetCvuCvuRoutingInfo()
                assert.True(t, ok, "Debtor RoutingInfo should be CvuCvuRoutingInfo")
                assert.Equal(t, "cvu", debtorCvuRoutingInfo.CvuRoutingInfoType)
                assert.Equal(t, "1122334455667788554433", debtorCvuRoutingInfo.Cvu)

                // Beneficiary fields
                assert.True(t, paymentCreatedEvent.Data.Beneficiary.InstitutionName.IsSet(), "Beneficiary InstitutionName should be set")
                assert.Equal(t, "dinopay", paymentCreatedEvent.Data.Beneficiary.InstitutionName.Value)
                assert.True(t, paymentCreatedEvent.Data.Beneficiary.InstitutionId.IsSet(), "Beneficiary InstitutionId should be set")
                assert.Equal(t, "dinopay", paymentCreatedEvent.Data.Beneficiary.InstitutionId.Value)
                assert.Equal(t, "ARS", string(paymentCreatedEvent.Data.Beneficiary.Currency))

                // Check if Beneficiary AccountDetails is CvuAccountDetails
                beneficiaryCvuDetails, ok := paymentCreatedEvent.Data.Beneficiary.AccountDetails.OneOf.GetCvuAccountDetails()
                assert.True(t, ok, "Beneficiary AccountDetails should be CvuAccountDetails")
                assert.Equal(t, "cvu", beneficiaryCvuDetails.AccountType)
                assert.True(t, beneficiaryCvuDetails.Cuit.IsSet(), "Beneficiary Cuit should be set")
                assert.Equal(t, "23679876453", beneficiaryCvuDetails.Cuit.Value)

                // Check if Beneficiary RoutingInfo is CvuCvuRoutingInfo
                beneficiaryCvuRoutingInfo, ok := beneficiaryCvuDetails.RoutingInfo.OneOf.GetCvuCvuRoutingInfo()
                assert.True(t, ok, "Beneficiary RoutingInfo should be CvuCvuRoutingInfo")
                assert.Equal(t, "cvu", beneficiaryCvuRoutingInfo.CvuRoutingInfoType)
                assert.Equal(t, "1122334455667788554433", beneficiaryCvuRoutingInfo.Cvu)
            },
        )

        err = event.Accept(ctx, mockHandler)
        require.NoError(t, err)

        mock.AssertExpectationsForObjects(t, mockHandler)
    })

    t.Run("PaymentCreated with optional fields empty", func(t *testing.T) {
        jsonPayload, err := os.ReadFile("testdata/payment_created_dinopay_optional_empty.json")
        require.NoError(t, err)

        event, err := deserializer.Deserialize(jsonPayload)
        require.NoError(t, err)

        ctx := context.TODO()

        mockHandler := NewMockHandler(t)
        mockHandler.EXPECT().HandlePaymentCreated(ctx, event).Return(nil).Run(
            func(ctx context.Context, paymentCreatedEvent PaymentCreated) {
                // Debtor fields
                assert.False(t, paymentCreatedEvent.Data.Debtor.InstitutionName.IsSet(), "Debtor InstitutionName should be set")
                assert.False(t, paymentCreatedEvent.Data.Debtor.InstitutionId.IsSet(), "Debtor InstitutionId should be set")
                // Beneficiary fields
                assert.False(t, paymentCreatedEvent.Data.Beneficiary.InstitutionName.IsSet(), "Beneficiary InstitutionName should not be set")
                assert.False(t, paymentCreatedEvent.Data.Beneficiary.InstitutionId.IsSet(), "Beneficiary InstitutionId should not be set")

            },
        )

        err = event.Accept(ctx, mockHandler)
        require.NoError(t, err)

        mock.AssertExpectationsForObjects(t, mockHandler)
    })
}
