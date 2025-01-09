package tests

import (
    "context"
    "net/http/httptest"
    "testing"

    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/walletera/payments-types/api"
)

func TestClient_GetPayment(t *testing.T) {

    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    getPaymentParams := api.GetPaymentParams{PaymentId: paymentId}

    handlerMock := NewMockHandler(t)
    handlerMock.EXPECT().
        GetPayment(mock.Anything, getPaymentParams).
        Return(&api.Payment{
            ID: paymentId,
        }, nil)

    bearerAuth := api.BearerAuth{Token: "json.web.token"}

    securityHandlerMock := NewMockSecurityHandler(t)
    securityHandlerMock.EXPECT().
        HandleBearerAuth(
            mock.Anything,
            mock.Anything,
            mock.Anything,
        ).
        Return(context.Background(), nil).
        Run(func(ctx context.Context, operationName string, bearer api.BearerAuth) {
            assert.Equal(t, bearerAuth, bearer)
        })

    paymentsServer, err := api.NewServer(handlerMock, securityHandlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    securitySourceMock := NewMockSecuritySource(t)
    securitySourceMock.EXPECT().
        BearerAuth(mock.Anything, mock.Anything).
        Return(
            bearerAuth,
            nil,
        )

    paymentsClient, err := api.NewClient(ts.URL, securitySourceMock)
    require.NoError(t, err)

    resp, err := paymentsClient.GetPayment(context.Background(), getPaymentParams)
    require.NoError(t, err)

    payment, ok := resp.(*api.Payment)
    require.True(t, ok)

    require.Equal(t, paymentId, payment.ID)
}

func TestClient_PatchPayment(t *testing.T) {
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

    handlerMock := NewMockHandler(t)
    handlerMock.EXPECT().
        PatchPayment(mock.Anything, paymentPatchBody, patchWithdrawalParams).
        Return(&api.PatchPaymentOK{}, nil)

    bearerAuth := api.BearerAuth{Token: "json.web.token"}

    securityHandlerMock := NewMockSecurityHandler(t)
    securityHandlerMock.EXPECT().
        HandleBearerAuth(
            mock.Anything,
            mock.Anything,
            mock.Anything,
        ).
        Return(context.Background(), nil).
        Run(func(ctx context.Context, operationName string, bearer api.BearerAuth) {
            assert.Equal(t, bearerAuth, bearer)
        })

    paymentsServer, err := api.NewServer(handlerMock, securityHandlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    securitySourceMock := NewMockSecuritySource(t)
    securitySourceMock.EXPECT().
        BearerAuth(mock.Anything, mock.Anything).
        Return(
            bearerAuth,
            nil,
        )

    paymentsClient, err := api.NewClient(ts.URL, securitySourceMock)
    require.NoError(t, err)

    _, err = paymentsClient.PatchPayment(context.Background(), paymentPatchBody, patchWithdrawalParams)

    require.NoError(t, err)
}

func TestClient_PostPayment(t *testing.T) {
    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    customerId, err := uuid.NewUUID()
    require.NoError(t, err)

    externalId, err := uuid.NewUUID()
    require.NoError(t, err)

    payment := &api.Payment{
        ID:       paymentId,
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

    handlerMock := NewMockHandler(t)
    handlerMock.EXPECT().
        PostPayment(mock.Anything, payment, postParams).
        Return(payment, nil)

    bearerAuth := api.BearerAuth{Token: "json.web.token"}

    securityHandlerMock := NewMockSecurityHandler(t)
    securityHandlerMock.EXPECT().
        HandleBearerAuth(
            mock.Anything,
            mock.Anything,
            mock.Anything,
        ).
        Return(context.Background(), nil).
        Run(func(ctx context.Context, operationName string, bearer api.BearerAuth) {
            assert.Equal(t, bearerAuth, bearer)
        })

    paymentsServer, err := api.NewServer(handlerMock, securityHandlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    securitySourceMock := NewMockSecuritySource(t)
    securitySourceMock.EXPECT().
        BearerAuth(mock.Anything, mock.Anything).
        Return(
            bearerAuth,
            nil,
        )

    paymentsClient, err := api.NewClient(ts.URL, securitySourceMock)
    require.NoError(t, err)

    resp, err := paymentsClient.PostPayment(context.Background(), payment, postParams)
    require.NoError(t, err)

    _, ok := resp.(*api.Payment)
    require.True(t, ok)
}
