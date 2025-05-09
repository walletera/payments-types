package privateapi

import (
    "context"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    privapiBuilder "github.com/walletera/payments-types/builders/privateapi"
    "github.com/walletera/payments-types/privateapi"
    "github.com/walletera/payments-types/tests/mocks"
)

func TestClient_GetPayment(t *testing.T) {

    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    debtor := debtor()
    beneficiary := beneficiary()

    getPaymentParams := privateapi.GetPaymentParams{PaymentId: paymentId}

    handlerMock := mocks.NewMockhandler_privateapi(t)
    handlerMock.EXPECT().
        GetPayment(mock.Anything, getPaymentParams).
        Return(&privateapi.Payment{
            ID:          paymentId,
            Amount:      0,
            Currency:    privateapi.CurrencyUSD,
            Gateway:     privateapi.GatewayBind,
            Debtor:      debtor,
            Beneficiary: beneficiary,
            Direction:   privateapi.DirectionOutbound,
            CustomerId:  uuid.UUID{},
            Status:      privateapi.PaymentStatusConfirmed,
            ExternalId:  privateapi.OptString{},
            SchemeId:    privateapi.OptString{},
            CreatedAt:   time.Time{},
            UpdatedAt:   time.Time{},
        }, nil)

    paymentsServer, err := privateapi.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    paymentsClient, err := privateapi.NewClient(ts.URL)
    require.NoError(t, err)

    resp, err := paymentsClient.GetPayment(context.Background(), getPaymentParams)
    require.NoError(t, err)

    payment, ok := resp.(*privateapi.Payment)
    require.True(t, ok)

    require.Equal(t, paymentId, payment.ID)
}

func TestClient_PatchPayment(t *testing.T) {
    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    externalId, err := uuid.NewUUID()
    require.NoError(t, err)

    paymentPatchBody := &privateapi.PaymentUpdate{
        ExternalId: privateapi.OptString{
            Value: externalId.String(),
            Set:   true,
        },
        Status: privateapi.PaymentStatusConfirmed,
    }

    patchWithdrawalParams := privateapi.PatchPaymentParams{
        PaymentId: paymentId,
    }

    handlerMock := mocks.NewMockhandler_privateapi(t)
    handlerMock.EXPECT().
        PatchPayment(mock.Anything, paymentPatchBody, patchWithdrawalParams).
        Return(&privateapi.PatchPaymentOK{}, nil)

    paymentsServer, err := privateapi.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    paymentsClient, err := privateapi.NewClient(ts.URL)
    require.NoError(t, err)

    _, err = paymentsClient.PatchPayment(context.Background(), paymentPatchBody, patchWithdrawalParams)

    require.NoError(t, err)
}

func TestClient_PostPayment(t *testing.T) {
    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    debtor := debtor()
    beneficiary := beneficiary()

    paymentCreationReq := &privateapi.PostPaymentReq{
        ID:          paymentId,
        Amount:      100,
        Currency:    "USD",
        Gateway:     "bind",
        Direction:   privateapi.DirectionOutbound,
        Status:      privateapi.PaymentStatusConfirmed,
        Debtor:      debtor,
        Beneficiary: beneficiary,
    }

    correlationID, err := uuid.NewUUID()
    require.NoError(t, err)

    postParams := privateapi.PostPaymentParams{
        XWalleteraCorrelationID: privateapi.NewOptUUID(correlationID),
    }

    handlerMock := mocks.NewMockhandler_privateapi(t)
    handlerMock.EXPECT().
        PostPayment(mock.Anything, paymentCreationReq, postParams).
        Return(&privateapi.Payment{
            ID:          paymentId,
            Amount:      0,
            Currency:    "USD",
            Debtor:      debtor,
            Beneficiary: beneficiary,
            Direction:   privateapi.DirectionOutbound,
            CustomerId:  uuid.UUID{},
            ExternalId:  privateapi.OptString{},
            SchemeId:    privateapi.OptString{},
            Status:      privateapi.PaymentStatusConfirmed,
            Gateway:     "bind",
            CreatedAt:   time.Time{},
            UpdatedAt:   time.Time{},
        }, nil)

    paymentsServer, err := privateapi.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    paymentsClient, err := privateapi.NewClient(ts.URL)
    require.NoError(t, err)

    resp, err := paymentsClient.PostPayment(context.Background(), paymentCreationReq, postParams)
    require.NoError(t, err)

    _, ok := resp.(*privateapi.Payment)
    require.True(t, ok)
}

func beneficiary() privateapi.Account {
    account, err := privapiBuilder.NewCVUAccountBuilder().
        WithCVU("2222222222222222222222").
        Build()
    if err != nil {
        panic(err)
    }
    return account
}

func debtor() privateapi.Account {
    account, err := privapiBuilder.NewCVUAccountBuilder().
        WithCVU("1111111111111111111111").
        Build()
    if err != nil {
        panic(err)
    }
    return account
}
