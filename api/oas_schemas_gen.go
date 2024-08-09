// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// Bank account details.
// Ref: #/components/schemas/accountDetails
type AccountDetails struct {
	BankName          OptString `json:"bankName"`
	BankId            OptString `json:"bankId"`
	AccountHolder     OptString `json:"accountHolder"`
	AccountNumber     OptString `json:"accountNumber"`
	AccountNumberType OptString `json:"accountNumberType"`
	RoutingKey        OptString `json:"routingKey"`
}

// GetBankName returns the value of BankName.
func (s *AccountDetails) GetBankName() OptString {
	return s.BankName
}

// GetBankId returns the value of BankId.
func (s *AccountDetails) GetBankId() OptString {
	return s.BankId
}

// GetAccountHolder returns the value of AccountHolder.
func (s *AccountDetails) GetAccountHolder() OptString {
	return s.AccountHolder
}

// GetAccountNumber returns the value of AccountNumber.
func (s *AccountDetails) GetAccountNumber() OptString {
	return s.AccountNumber
}

// GetAccountNumberType returns the value of AccountNumberType.
func (s *AccountDetails) GetAccountNumberType() OptString {
	return s.AccountNumberType
}

// GetRoutingKey returns the value of RoutingKey.
func (s *AccountDetails) GetRoutingKey() OptString {
	return s.RoutingKey
}

// SetBankName sets the value of BankName.
func (s *AccountDetails) SetBankName(val OptString) {
	s.BankName = val
}

// SetBankId sets the value of BankId.
func (s *AccountDetails) SetBankId(val OptString) {
	s.BankId = val
}

// SetAccountHolder sets the value of AccountHolder.
func (s *AccountDetails) SetAccountHolder(val OptString) {
	s.AccountHolder = val
}

// SetAccountNumber sets the value of AccountNumber.
func (s *AccountDetails) SetAccountNumber(val OptString) {
	s.AccountNumber = val
}

// SetAccountNumberType sets the value of AccountNumberType.
func (s *AccountDetails) SetAccountNumberType(val OptString) {
	s.AccountNumberType = val
}

// SetRoutingKey sets the value of RoutingKey.
func (s *AccountDetails) SetRoutingKey(val OptString) {
	s.RoutingKey = val
}

type ErrorMessage string

func (*ErrorMessage) postPaymentRes() {}

// NewOptAccountDetails returns new OptAccountDetails with value set to v.
func NewOptAccountDetails(v AccountDetails) OptAccountDetails {
	return OptAccountDetails{
		Value: v,
		Set:   true,
	}
}

// OptAccountDetails is optional AccountDetails.
type OptAccountDetails struct {
	Value AccountDetails
	Set   bool
}

// IsSet returns true if OptAccountDetails was set.
func (o OptAccountDetails) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAccountDetails) Reset() {
	var v AccountDetails
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAccountDetails) SetTo(v AccountDetails) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAccountDetails) Get() (v AccountDetails, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAccountDetails) Or(d AccountDetails) AccountDetails {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPaymentStatus returns new OptPaymentStatus with value set to v.
func NewOptPaymentStatus(v PaymentStatus) OptPaymentStatus {
	return OptPaymentStatus{
		Value: v,
		Set:   true,
	}
}

// OptPaymentStatus is optional PaymentStatus.
type OptPaymentStatus struct {
	Value PaymentStatus
	Set   bool
}

// IsSet returns true if OptPaymentStatus was set.
func (o OptPaymentStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPaymentStatus) Reset() {
	var v PaymentStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPaymentStatus) SetTo(v PaymentStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPaymentStatus) Get() (v PaymentStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPaymentStatus) Or(d PaymentStatus) PaymentStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUUID returns new OptUUID with value set to v.
func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}

// OptUUID is optional uuid.UUID.
type OptUUID struct {
	Value uuid.UUID
	Set   bool
}

// IsSet returns true if OptUUID was set.
func (o OptUUID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUUID) Or(d uuid.UUID) uuid.UUID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// PatchPaymentOK is response for PatchPayment operation.
type PatchPaymentOK struct{}

func (*PatchPaymentOK) patchPaymentRes() {}

// Payment type.
// Ref: #/components/schemas/payment
type Payment struct {
	// Payment id.
	ID uuid.UUID `json:"id"`
	// Payment amount.
	Amount float64 `json:"amount"`
	// Payment currency.
	Currency string `json:"currency"`
	// The customer that sent or received the payment.
	CustomerId OptUUID `json:"customerId"`
	// Id assigned to the operation by an external payment provider.
	ExternalId  OptUUID           `json:"externalId"`
	Beneficiary OptAccountDetails `json:"beneficiary"`
	Debtor      OptAccountDetails `json:"debtor"`
	Status      OptPaymentStatus  `json:"status"`
	CreatedAt   OptDateTime       `json:"createdAt"`
	UpdatedAt   OptDateTime       `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Payment) GetID() uuid.UUID {
	return s.ID
}

// GetAmount returns the value of Amount.
func (s *Payment) GetAmount() float64 {
	return s.Amount
}

// GetCurrency returns the value of Currency.
func (s *Payment) GetCurrency() string {
	return s.Currency
}

// GetCustomerId returns the value of CustomerId.
func (s *Payment) GetCustomerId() OptUUID {
	return s.CustomerId
}

// GetExternalId returns the value of ExternalId.
func (s *Payment) GetExternalId() OptUUID {
	return s.ExternalId
}

// GetBeneficiary returns the value of Beneficiary.
func (s *Payment) GetBeneficiary() OptAccountDetails {
	return s.Beneficiary
}

// GetDebtor returns the value of Debtor.
func (s *Payment) GetDebtor() OptAccountDetails {
	return s.Debtor
}

// GetStatus returns the value of Status.
func (s *Payment) GetStatus() OptPaymentStatus {
	return s.Status
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Payment) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Payment) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Payment) SetID(val uuid.UUID) {
	s.ID = val
}

// SetAmount sets the value of Amount.
func (s *Payment) SetAmount(val float64) {
	s.Amount = val
}

// SetCurrency sets the value of Currency.
func (s *Payment) SetCurrency(val string) {
	s.Currency = val
}

// SetCustomerId sets the value of CustomerId.
func (s *Payment) SetCustomerId(val OptUUID) {
	s.CustomerId = val
}

// SetExternalId sets the value of ExternalId.
func (s *Payment) SetExternalId(val OptUUID) {
	s.ExternalId = val
}

// SetBeneficiary sets the value of Beneficiary.
func (s *Payment) SetBeneficiary(val OptAccountDetails) {
	s.Beneficiary = val
}

// SetDebtor sets the value of Debtor.
func (s *Payment) SetDebtor(val OptAccountDetails) {
	s.Debtor = val
}

// SetStatus sets the value of Status.
func (s *Payment) SetStatus(val OptPaymentStatus) {
	s.Status = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Payment) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Payment) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

func (*Payment) patchPaymentRes() {}

// Body of the PATH /withdrawal request.
// Ref: #/components/schemas/paymentPatchBody
type PaymentPatchBody struct {
	// Id assigned to the operation by the external payment provider.
	ExternalId OptUUID `json:"externalId"`
	// Withdrawal status.
	Status PaymentPatchBodyStatus `json:"status"`
}

// GetExternalId returns the value of ExternalId.
func (s *PaymentPatchBody) GetExternalId() OptUUID {
	return s.ExternalId
}

// GetStatus returns the value of Status.
func (s *PaymentPatchBody) GetStatus() PaymentPatchBodyStatus {
	return s.Status
}

// SetExternalId sets the value of ExternalId.
func (s *PaymentPatchBody) SetExternalId(val OptUUID) {
	s.ExternalId = val
}

// SetStatus sets the value of Status.
func (s *PaymentPatchBody) SetStatus(val PaymentPatchBodyStatus) {
	s.Status = val
}

// Withdrawal status.
type PaymentPatchBodyStatus string

const (
	PaymentPatchBodyStatusPending   PaymentPatchBodyStatus = "pending"
	PaymentPatchBodyStatusConfirmed PaymentPatchBodyStatus = "confirmed"
	PaymentPatchBodyStatusRejected  PaymentPatchBodyStatus = "rejected"
)

// AllValues returns all PaymentPatchBodyStatus values.
func (PaymentPatchBodyStatus) AllValues() []PaymentPatchBodyStatus {
	return []PaymentPatchBodyStatus{
		PaymentPatchBodyStatusPending,
		PaymentPatchBodyStatusConfirmed,
		PaymentPatchBodyStatusRejected,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s PaymentPatchBodyStatus) MarshalText() ([]byte, error) {
	switch s {
	case PaymentPatchBodyStatusPending:
		return []byte(s), nil
	case PaymentPatchBodyStatusConfirmed:
		return []byte(s), nil
	case PaymentPatchBodyStatusRejected:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *PaymentPatchBodyStatus) UnmarshalText(data []byte) error {
	switch PaymentPatchBodyStatus(data) {
	case PaymentPatchBodyStatusPending:
		*s = PaymentPatchBodyStatusPending
		return nil
	case PaymentPatchBodyStatusConfirmed:
		*s = PaymentPatchBodyStatusConfirmed
		return nil
	case PaymentPatchBodyStatusRejected:
		*s = PaymentPatchBodyStatusRejected
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusDelivered PaymentStatus = "delivered"
	PaymentStatusConfirmed PaymentStatus = "confirmed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

// AllValues returns all PaymentStatus values.
func (PaymentStatus) AllValues() []PaymentStatus {
	return []PaymentStatus{
		PaymentStatusPending,
		PaymentStatusDelivered,
		PaymentStatusConfirmed,
		PaymentStatusFailed,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s PaymentStatus) MarshalText() ([]byte, error) {
	switch s {
	case PaymentStatusPending:
		return []byte(s), nil
	case PaymentStatusDelivered:
		return []byte(s), nil
	case PaymentStatusConfirmed:
		return []byte(s), nil
	case PaymentStatusFailed:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *PaymentStatus) UnmarshalText(data []byte) error {
	switch PaymentStatus(data) {
	case PaymentStatusPending:
		*s = PaymentStatusPending
		return nil
	case PaymentStatusDelivered:
		*s = PaymentStatusDelivered
		return nil
	case PaymentStatusConfirmed:
		*s = PaymentStatusConfirmed
		return nil
	case PaymentStatusFailed:
		*s = PaymentStatusFailed
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// PostPaymentCreated is response for PostPayment operation.
type PostPaymentCreated struct{}

func (*PostPaymentCreated) postPaymentRes() {}