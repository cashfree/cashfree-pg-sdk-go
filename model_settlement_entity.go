/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
)

// checks if the SettlementEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementEntity{}

// SettlementEntity Settlement entity object
type SettlementEntity struct {
	CfPaymentId *int64 `json:"cf_payment_id,omitempty"`
	CfSettlementId *int64 `json:"cf_settlement_id,omitempty"`
	SettlementCurrency *string `json:"settlement_currency,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	Entity *string `json:"entity,omitempty"`
	OrderAmount *float32 `json:"order_amount,omitempty"`
	PaymentTime *string `json:"payment_time,omitempty"`
	ServiceCharge *float32 `json:"service_charge,omitempty"`
	ServiceTax *float32 `json:"service_tax,omitempty"`
	SettlementAmount *float32 `json:"settlement_amount,omitempty"`
	SettlementId *int64 `json:"settlement_id,omitempty"`
	TransferId *int64 `json:"transfer_id,omitempty"`
	TransferTime *string `json:"transfer_time,omitempty"`
	TransferUtr *string `json:"transfer_utr,omitempty"`
}

// NewSettlementEntity instantiates a new SettlementEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettlementEntity() *SettlementEntity {
	this := SettlementEntity{}
	return &this
}

// NewSettlementEntityWithDefaults instantiates a new SettlementEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettlementEntityWithDefaults() *SettlementEntity {
	this := SettlementEntity{}
	return &this
}

// GetCfPaymentId returns the CfPaymentId field value if set, zero value otherwise.
func (o *SettlementEntity) GetCfPaymentId() int64 {
	if o == nil || IsNil(o.CfPaymentId) {
		var ret int64
		return ret
	}
	return *o.CfPaymentId
}

// GetCfPaymentIdOk returns a tuple with the CfPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetCfPaymentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CfPaymentId) {
		return nil, false
	}
	return o.CfPaymentId, true
}

// HasCfPaymentId returns a boolean if a field has been set.
func (o *SettlementEntity) HasCfPaymentId() bool {
	if o != nil && !IsNil(o.CfPaymentId) {
		return true
	}

	return false
}

// SetCfPaymentId gets a reference to the given int64 and assigns it to the CfPaymentId field.
func (o *SettlementEntity) SetCfPaymentId(v int64) {
	o.CfPaymentId = &v
}

// GetCfSettlementId returns the CfSettlementId field value if set, zero value otherwise.
func (o *SettlementEntity) GetCfSettlementId() int64 {
	if o == nil || IsNil(o.CfSettlementId) {
		var ret int64
		return ret
	}
	return *o.CfSettlementId
}

// GetCfSettlementIdOk returns a tuple with the CfSettlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetCfSettlementIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CfSettlementId) {
		return nil, false
	}
	return o.CfSettlementId, true
}

// HasCfSettlementId returns a boolean if a field has been set.
func (o *SettlementEntity) HasCfSettlementId() bool {
	if o != nil && !IsNil(o.CfSettlementId) {
		return true
	}

	return false
}

// SetCfSettlementId gets a reference to the given int64 and assigns it to the CfSettlementId field.
func (o *SettlementEntity) SetCfSettlementId(v int64) {
	o.CfSettlementId = &v
}

// GetSettlementCurrency returns the SettlementCurrency field value if set, zero value otherwise.
func (o *SettlementEntity) GetSettlementCurrency() string {
	if o == nil || IsNil(o.SettlementCurrency) {
		var ret string
		return ret
	}
	return *o.SettlementCurrency
}

// GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetSettlementCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SettlementCurrency) {
		return nil, false
	}
	return o.SettlementCurrency, true
}

// HasSettlementCurrency returns a boolean if a field has been set.
func (o *SettlementEntity) HasSettlementCurrency() bool {
	if o != nil && !IsNil(o.SettlementCurrency) {
		return true
	}

	return false
}

// SetSettlementCurrency gets a reference to the given string and assigns it to the SettlementCurrency field.
func (o *SettlementEntity) SetSettlementCurrency(v string) {
	o.SettlementCurrency = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SettlementEntity) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SettlementEntity) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *SettlementEntity) SetOrderId(v string) {
	o.OrderId = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *SettlementEntity) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *SettlementEntity) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *SettlementEntity) SetEntity(v string) {
	o.Entity = &v
}

// GetOrderAmount returns the OrderAmount field value if set, zero value otherwise.
func (o *SettlementEntity) GetOrderAmount() float32 {
	if o == nil || IsNil(o.OrderAmount) {
		var ret float32
		return ret
	}
	return *o.OrderAmount
}

// GetOrderAmountOk returns a tuple with the OrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetOrderAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderAmount) {
		return nil, false
	}
	return o.OrderAmount, true
}

// HasOrderAmount returns a boolean if a field has been set.
func (o *SettlementEntity) HasOrderAmount() bool {
	if o != nil && !IsNil(o.OrderAmount) {
		return true
	}

	return false
}

// SetOrderAmount gets a reference to the given float32 and assigns it to the OrderAmount field.
func (o *SettlementEntity) SetOrderAmount(v float32) {
	o.OrderAmount = &v
}

// GetPaymentTime returns the PaymentTime field value if set, zero value otherwise.
func (o *SettlementEntity) GetPaymentTime() string {
	if o == nil || IsNil(o.PaymentTime) {
		var ret string
		return ret
	}
	return *o.PaymentTime
}

// GetPaymentTimeOk returns a tuple with the PaymentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetPaymentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTime) {
		return nil, false
	}
	return o.PaymentTime, true
}

// HasPaymentTime returns a boolean if a field has been set.
func (o *SettlementEntity) HasPaymentTime() bool {
	if o != nil && !IsNil(o.PaymentTime) {
		return true
	}

	return false
}

// SetPaymentTime gets a reference to the given string and assigns it to the PaymentTime field.
func (o *SettlementEntity) SetPaymentTime(v string) {
	o.PaymentTime = &v
}

// GetServiceCharge returns the ServiceCharge field value if set, zero value otherwise.
func (o *SettlementEntity) GetServiceCharge() float32 {
	if o == nil || IsNil(o.ServiceCharge) {
		var ret float32
		return ret
	}
	return *o.ServiceCharge
}

// GetServiceChargeOk returns a tuple with the ServiceCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetServiceChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.ServiceCharge) {
		return nil, false
	}
	return o.ServiceCharge, true
}

// HasServiceCharge returns a boolean if a field has been set.
func (o *SettlementEntity) HasServiceCharge() bool {
	if o != nil && !IsNil(o.ServiceCharge) {
		return true
	}

	return false
}

// SetServiceCharge gets a reference to the given float32 and assigns it to the ServiceCharge field.
func (o *SettlementEntity) SetServiceCharge(v float32) {
	o.ServiceCharge = &v
}

// GetServiceTax returns the ServiceTax field value if set, zero value otherwise.
func (o *SettlementEntity) GetServiceTax() float32 {
	if o == nil || IsNil(o.ServiceTax) {
		var ret float32
		return ret
	}
	return *o.ServiceTax
}

// GetServiceTaxOk returns a tuple with the ServiceTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetServiceTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.ServiceTax) {
		return nil, false
	}
	return o.ServiceTax, true
}

// HasServiceTax returns a boolean if a field has been set.
func (o *SettlementEntity) HasServiceTax() bool {
	if o != nil && !IsNil(o.ServiceTax) {
		return true
	}

	return false
}

// SetServiceTax gets a reference to the given float32 and assigns it to the ServiceTax field.
func (o *SettlementEntity) SetServiceTax(v float32) {
	o.ServiceTax = &v
}

// GetSettlementAmount returns the SettlementAmount field value if set, zero value otherwise.
func (o *SettlementEntity) GetSettlementAmount() float32 {
	if o == nil || IsNil(o.SettlementAmount) {
		var ret float32
		return ret
	}
	return *o.SettlementAmount
}

// GetSettlementAmountOk returns a tuple with the SettlementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetSettlementAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.SettlementAmount) {
		return nil, false
	}
	return o.SettlementAmount, true
}

// HasSettlementAmount returns a boolean if a field has been set.
func (o *SettlementEntity) HasSettlementAmount() bool {
	if o != nil && !IsNil(o.SettlementAmount) {
		return true
	}

	return false
}

// SetSettlementAmount gets a reference to the given float32 and assigns it to the SettlementAmount field.
func (o *SettlementEntity) SetSettlementAmount(v float32) {
	o.SettlementAmount = &v
}

// GetSettlementId returns the SettlementId field value if set, zero value otherwise.
func (o *SettlementEntity) GetSettlementId() int64 {
	if o == nil || IsNil(o.SettlementId) {
		var ret int64
		return ret
	}
	return *o.SettlementId
}

// GetSettlementIdOk returns a tuple with the SettlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetSettlementIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SettlementId) {
		return nil, false
	}
	return o.SettlementId, true
}

// HasSettlementId returns a boolean if a field has been set.
func (o *SettlementEntity) HasSettlementId() bool {
	if o != nil && !IsNil(o.SettlementId) {
		return true
	}

	return false
}

// SetSettlementId gets a reference to the given int64 and assigns it to the SettlementId field.
func (o *SettlementEntity) SetSettlementId(v int64) {
	o.SettlementId = &v
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *SettlementEntity) GetTransferId() int64 {
	if o == nil || IsNil(o.TransferId) {
		var ret int64
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetTransferIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *SettlementEntity) HasTransferId() bool {
	if o != nil && !IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given int64 and assigns it to the TransferId field.
func (o *SettlementEntity) SetTransferId(v int64) {
	o.TransferId = &v
}

// GetTransferTime returns the TransferTime field value if set, zero value otherwise.
func (o *SettlementEntity) GetTransferTime() string {
	if o == nil || IsNil(o.TransferTime) {
		var ret string
		return ret
	}
	return *o.TransferTime
}

// GetTransferTimeOk returns a tuple with the TransferTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetTransferTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TransferTime) {
		return nil, false
	}
	return o.TransferTime, true
}

// HasTransferTime returns a boolean if a field has been set.
func (o *SettlementEntity) HasTransferTime() bool {
	if o != nil && !IsNil(o.TransferTime) {
		return true
	}

	return false
}

// SetTransferTime gets a reference to the given string and assigns it to the TransferTime field.
func (o *SettlementEntity) SetTransferTime(v string) {
	o.TransferTime = &v
}

// GetTransferUtr returns the TransferUtr field value if set, zero value otherwise.
func (o *SettlementEntity) GetTransferUtr() string {
	if o == nil || IsNil(o.TransferUtr) {
		var ret string
		return ret
	}
	return *o.TransferUtr
}

// GetTransferUtrOk returns a tuple with the TransferUtr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementEntity) GetTransferUtrOk() (*string, bool) {
	if o == nil || IsNil(o.TransferUtr) {
		return nil, false
	}
	return o.TransferUtr, true
}

// HasTransferUtr returns a boolean if a field has been set.
func (o *SettlementEntity) HasTransferUtr() bool {
	if o != nil && !IsNil(o.TransferUtr) {
		return true
	}

	return false
}

// SetTransferUtr gets a reference to the given string and assigns it to the TransferUtr field.
func (o *SettlementEntity) SetTransferUtr(v string) {
	o.TransferUtr = &v
}

func (o SettlementEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.CfSettlementId) {
		toSerialize["cf_settlement_id"] = o.CfSettlementId
	}
	if !IsNil(o.SettlementCurrency) {
		toSerialize["settlement_currency"] = o.SettlementCurrency
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.OrderAmount) {
		toSerialize["order_amount"] = o.OrderAmount
	}
	if !IsNil(o.PaymentTime) {
		toSerialize["payment_time"] = o.PaymentTime
	}
	if !IsNil(o.ServiceCharge) {
		toSerialize["service_charge"] = o.ServiceCharge
	}
	if !IsNil(o.ServiceTax) {
		toSerialize["service_tax"] = o.ServiceTax
	}
	if !IsNil(o.SettlementAmount) {
		toSerialize["settlement_amount"] = o.SettlementAmount
	}
	if !IsNil(o.SettlementId) {
		toSerialize["settlement_id"] = o.SettlementId
	}
	if !IsNil(o.TransferId) {
		toSerialize["transfer_id"] = o.TransferId
	}
	if !IsNil(o.TransferTime) {
		toSerialize["transfer_time"] = o.TransferTime
	}
	if !IsNil(o.TransferUtr) {
		toSerialize["transfer_utr"] = o.TransferUtr
	}
	return toSerialize, nil
}

type NullableSettlementEntity struct {
	value *SettlementEntity
	isSet bool
}

func (v NullableSettlementEntity) Get() *SettlementEntity {
	return v.value
}

func (v *NullableSettlementEntity) Set(val *SettlementEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableSettlementEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableSettlementEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettlementEntity(val *SettlementEntity) *NullableSettlementEntity {
	return &NullableSettlementEntity{value: val, isSet: true}
}

func (v NullableSettlementEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettlementEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


