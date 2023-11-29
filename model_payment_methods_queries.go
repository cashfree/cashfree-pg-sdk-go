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

// checks if the PaymentMethodsQueries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsQueries{}

// PaymentMethodsQueries Payment Method Query Object
type PaymentMethodsQueries struct {
	// Amount of the order.
	Amount *float32 `json:"amount,omitempty"`
	// OrderId of the order. Either of `order_id` or `order_amount` is mandatory.
	OrderId *string `json:"order_id,omitempty"`
}

// NewPaymentMethodsQueries instantiates a new PaymentMethodsQueries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsQueries() *PaymentMethodsQueries {
	this := PaymentMethodsQueries{}
	return &this
}

// NewPaymentMethodsQueriesWithDefaults instantiates a new PaymentMethodsQueries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsQueriesWithDefaults() *PaymentMethodsQueries {
	this := PaymentMethodsQueries{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentMethodsQueries) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsQueries) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentMethodsQueries) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PaymentMethodsQueries) SetAmount(v float32) {
	o.Amount = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PaymentMethodsQueries) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsQueries) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PaymentMethodsQueries) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *PaymentMethodsQueries) SetOrderId(v string) {
	o.OrderId = &v
}

func (o PaymentMethodsQueries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsQueries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	return toSerialize, nil
}

type NullablePaymentMethodsQueries struct {
	value *PaymentMethodsQueries
	isSet bool
}

func (v NullablePaymentMethodsQueries) Get() *PaymentMethodsQueries {
	return v.value
}

func (v *NullablePaymentMethodsQueries) Set(val *PaymentMethodsQueries) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsQueries) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsQueries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsQueries(val *PaymentMethodsQueries) *NullablePaymentMethodsQueries {
	return &NullablePaymentMethodsQueries{value: val, isSet: true}
}

func (v NullablePaymentMethodsQueries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsQueries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


