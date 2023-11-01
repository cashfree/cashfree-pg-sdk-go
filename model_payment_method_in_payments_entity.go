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

// checks if the PaymentMethodInPaymentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodInPaymentsEntity{}

// PaymentMethodInPaymentsEntity payment methods all
type PaymentMethodInPaymentsEntity struct {
	PaymentMethod *PaymentMethodInPaymentsEntityPaymentMethod `json:"payment_method,omitempty"`
}

// NewPaymentMethodInPaymentsEntity instantiates a new PaymentMethodInPaymentsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodInPaymentsEntity() *PaymentMethodInPaymentsEntity {
	this := PaymentMethodInPaymentsEntity{}
	return &this
}

// NewPaymentMethodInPaymentsEntityWithDefaults instantiates a new PaymentMethodInPaymentsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodInPaymentsEntityWithDefaults() *PaymentMethodInPaymentsEntity {
	this := PaymentMethodInPaymentsEntity{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentMethodInPaymentsEntity) GetPaymentMethod() PaymentMethodInPaymentsEntityPaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret PaymentMethodInPaymentsEntityPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodInPaymentsEntity) GetPaymentMethodOk() (*PaymentMethodInPaymentsEntityPaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentMethodInPaymentsEntity) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethodInPaymentsEntityPaymentMethod and assigns it to the PaymentMethod field.
func (o *PaymentMethodInPaymentsEntity) SetPaymentMethod(v PaymentMethodInPaymentsEntityPaymentMethod) {
	o.PaymentMethod = &v
}

func (o PaymentMethodInPaymentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodInPaymentsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return toSerialize, nil
}

type NullablePaymentMethodInPaymentsEntity struct {
	value *PaymentMethodInPaymentsEntity
	isSet bool
}

func (v NullablePaymentMethodInPaymentsEntity) Get() *PaymentMethodInPaymentsEntity {
	return v.value
}

func (v *NullablePaymentMethodInPaymentsEntity) Set(val *PaymentMethodInPaymentsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodInPaymentsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodInPaymentsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodInPaymentsEntity(val *PaymentMethodInPaymentsEntity) *NullablePaymentMethodInPaymentsEntity {
	return &NullablePaymentMethodInPaymentsEntity{value: val, isSet: true}
}

func (v NullablePaymentMethodInPaymentsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodInPaymentsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

