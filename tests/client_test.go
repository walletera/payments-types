package tests

import (
    "context"
    "net/http/httptest"
    "testing"

    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/walletera/payments-types/api"
)

func TestClient_PatchWithdrawal(t *testing.T) {
    handlerMock := NewMockHandler(t)

    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    externalId, err := uuid.NewUUID()
    require.NoError(t, err)

    paymentPatchBody := &api.PaymentUpdate{
        ExternalId: api.OptUUID{
            Value: externalId,
            Set:   true,
        },
        Status: api.PaymentStatusConfirmed,
    }

    patchWithdrawalParams := api.PatchPaymentParams{
        PaymentId: paymentId,
    }

    handlerMock.EXPECT().
        PatchPayment(mock.Anything, paymentPatchBody, patchWithdrawalParams).
        Return(&api.PatchPaymentOK{}, nil)

    paymentsServer, err := api.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    paymentsClient, err := api.NewClient(ts.URL)
    require.NoError(t, err)

    _, err = paymentsClient.PatchPayment(context.Background(), paymentPatchBody, patchWithdrawalParams)

    require.NoError(t, err)
}

func TestClient_PostDeposit(t *testing.T) {
    handlerMock := NewMockHandler(t)

    depositlId, err := uuid.NewUUID()
    require.NoError(t, err)

    customerId, err := uuid.NewUUID()
    require.NoError(t, err)

    externalId, err := uuid.NewUUID()
    require.NoError(t, err)

    payment := &api.Payment{
        ID:       api.NewOptUUID(depositlId),
        Amount:   100,
        Currency: "usd",
        CustomerId: api.OptUUID{
            Value: customerId,
            Set:   true,
        },
        ExternalId: api.OptUUID{
            Value: externalId,
            Set:   true,
        },
    }

    correlationID, err := uuid.NewUUID()
    require.NoError(t, err)

    postParams := api.PostPaymentParams{
        XWalleteraCorrelationID: api.NewOptUUID(correlationID),
    }

    handlerMock.EXPECT().
        PostPayment(mock.Anything, payment, postParams).
        Return(payment, nil)

    paymentsServer, err := api.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    paymentsClient, err := api.NewClient(ts.URL)
    require.NoError(t, err)

    resp, err := paymentsClient.PostPayment(context.Background(), payment, postParams)
    require.NoError(t, err)

    _, ok := resp.(*api.Payment)
    require.True(t, ok)
}
