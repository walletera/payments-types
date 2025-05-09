package publicapi

import (
    "context"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/walletera/payments-types/builders/public"
    "github.com/walletera/payments-types/publicapi"
    "github.com/walletera/payments-types/tests/mocks"
)

func TestPublicAPI_GetPayment(t *testing.T) {

    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    debtor := debtor()
    beneficiary := beneficiary()

    getPaymentParams := publicapi.GetPaymentParams{PaymentId: paymentId}

    handlerMock := mocks.NewMockhandler_publicapi(t)
    handlerMock.EXPECT().
        GetPayment(mock.Anything, getPaymentParams).
        Return(&publicapi.Payment{
            ID:          paymentId,
            Amount:      0,
            Currency:    publicapi.CurrencyUSD,
            Gateway:     publicapi.GatewayBind,
            Debtor:      debtor,
            Beneficiary: beneficiary,
            Direction:   publicapi.DirectionOutbound,
            CustomerId:  uuid.UUID{},
            Status:      publicapi.PaymentStatusConfirmed,
            ExternalId:  publicapi.OptString{},
            SchemeId:    publicapi.OptString{},
            CreatedAt:   time.Time{},
            UpdatedAt:   time.Time{},
        }, nil)

    bearerAuth := publicapi.BearerAuth{Token: "json.web.token"}

    securityHandlerMock := mocks.NewMocksecurity_handler_publicapi(t)
    securityHandlerMock.EXPECT().
        HandleBearerAuth(
            mock.Anything,
            mock.Anything,
            mock.Anything,
        ).
        Return(context.Background(), nil).
        Run(func(ctx context.Context, operationName string, bearer publicapi.BearerAuth) {
            assert.Equal(t, bearerAuth, bearer)
        })

    paymentsServer, err := publicapi.NewServer(handlerMock, securityHandlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    securitySourceMock := mocks.NewMocksecurity_source_publicapi(t)
    securitySourceMock.EXPECT().
        BearerAuth(mock.Anything, mock.Anything).
        Return(
            bearerAuth,
            nil,
        )

    paymentsClient, err := publicapi.NewClient(ts.URL, securitySourceMock)
    require.NoError(t, err)

    resp, err := paymentsClient.GetPayment(context.Background(), getPaymentParams)
    require.NoError(t, err)

    payment, ok := resp.(*publicapi.Payment)
    require.True(t, ok)

    require.Equal(t, paymentId, payment.ID)
}

func TestPublicAPI_PostPayment(t *testing.T) {
    paymentId, err := uuid.NewUUID()
    require.NoError(t, err)

    debtor := debtor()
    beneficiary := beneficiary()

    paymentCreationReq := &publicapi.PostPaymentReq{
        ID:          paymentId,
        Amount:      100,
        Currency:    "USD",
        Gateway:     "bind",
        Debtor:      debtor,
        Beneficiary: beneficiary,
    }

    correlationID, err := uuid.NewUUID()
    require.NoError(t, err)

    postParams := publicapi.PostPaymentParams{
        XWalleteraCorrelationID: publicapi.NewOptUUID(correlationID),
    }

    handlerMock := mocks.NewMockhandler_publicapi(t)
    handlerMock.EXPECT().
        PostPayment(mock.Anything, paymentCreationReq, postParams).
        Return(&publicapi.Payment{
            ID:          paymentId,
            Amount:      0,
            Currency:    "USD",
            Debtor:      debtor,
            Beneficiary: beneficiary,
            Direction:   publicapi.DirectionOutbound,
            CustomerId:  uuid.UUID{},
            ExternalId:  publicapi.OptString{},
            SchemeId:    publicapi.OptString{},
            Status:      publicapi.PaymentStatusConfirmed,
            Gateway:     "bind",
            CreatedAt:   time.Time{},
            UpdatedAt:   time.Time{},
        }, nil)

    bearerAuth := publicapi.BearerAuth{Token: "json.web.token"}

    securityHandlerMock := mocks.NewMocksecurity_handler_publicapi(t)
    securityHandlerMock.EXPECT().
        HandleBearerAuth(
            mock.Anything,
            mock.Anything,
            mock.Anything,
        ).
        Return(context.Background(), nil).
        Run(func(ctx context.Context, operationName string, bearer publicapi.BearerAuth) {
            assert.Equal(t, bearerAuth, bearer)
        })

    paymentsServer, err := publicapi.NewServer(handlerMock, securityHandlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(paymentsServer)
    defer ts.Close()

    securitySourceMock := mocks.NewMocksecurity_source_publicapi(t)
    securitySourceMock.EXPECT().
        BearerAuth(mock.Anything, mock.Anything).
        Return(
            bearerAuth,
            nil,
        )

    paymentsClient, err := publicapi.NewClient(ts.URL, securitySourceMock)
    require.NoError(t, err)

    resp, err := paymentsClient.PostPayment(context.Background(), paymentCreationReq, postParams)
    require.NoError(t, err)

    _, ok := resp.(*publicapi.Payment)
    require.True(t, ok)
}

func beneficiary() publicapi.Account {
    account, err := public.NewCVUAccountBuilder().
        WithCVU("2222222222222222222222").
        Build()
    if err != nil {
        panic(err)
    }
    return account
}

func debtor() publicapi.Account {
    account, err := public.NewCVUAccountBuilder().
        WithCVU("1111111111111111111111").
        Build()
    if err != nil {
        panic(err)
    }
    return account
}
