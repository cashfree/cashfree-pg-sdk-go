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

// checks if the EligibilityPaymentMethodsEntityEntityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityPaymentMethodsEntityEntityDetails{}

// EligibilityPaymentMethodsEntityEntityDetails struct for EligibilityPaymentMethodsEntityEntityDetails
type EligibilityPaymentMethodsEntityEntityDetails struct {
	PaymentMethodDetails []PaymentModeDetails `json:"payment_method_details,omitempty"`
}

// NewEligibilityPaymentMethodsEntityEntityDetails instantiates a new EligibilityPaymentMethodsEntityEntityDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityPaymentMethodsEntityEntityDetails() *EligibilityPaymentMethodsEntityEntityDetails {
	this := EligibilityPaymentMethodsEntityEntityDetails{}
	return &this
}

// NewEligibilityPaymentMethodsEntityEntityDetailsWithDefaults instantiates a new EligibilityPaymentMethodsEntityEntityDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityPaymentMethodsEntityEntityDetailsWithDefaults() *EligibilityPaymentMethodsEntityEntityDetails {
	this := EligibilityPaymentMethodsEntityEntityDetails{}
	return &this
}

// GetPaymentMethodDetails returns the PaymentMethodDetails field value if set, zero value otherwise.
func (o *EligibilityPaymentMethodsEntityEntityDetails) GetPaymentMethodDetails() []PaymentModeDetails {
	if o == nil || IsNil(o.PaymentMethodDetails) {
		var ret []PaymentModeDetails
		return ret
	}
	return o.PaymentMethodDetails
}

// GetPaymentMethodDetailsOk returns a tuple with the PaymentMethodDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityPaymentMethodsEntityEntityDetails) GetPaymentMethodDetailsOk() ([]PaymentModeDetails, bool) {
	if o == nil || IsNil(o.PaymentMethodDetails) {
		return nil, false
	}
	return o.PaymentMethodDetails, true
}

// HasPaymentMethodDetails returns a boolean if a field has been set.
func (o *EligibilityPaymentMethodsEntityEntityDetails) HasPaymentMethodDetails() bool {
	if o != nil && !IsNil(o.PaymentMethodDetails) {
		return true
	}

	return false
}

// SetPaymentMethodDetails gets a reference to the given []PaymentModeDetails and assigns it to the PaymentMethodDetails field.
func (o *EligibilityPaymentMethodsEntityEntityDetails) SetPaymentMethodDetails(v []PaymentModeDetails) {
	o.PaymentMethodDetails = v
}

func (o EligibilityPaymentMethodsEntityEntityDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityPaymentMethodsEntityEntityDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethodDetails) {
		toSerialize["payment_method_details"] = o.PaymentMethodDetails
	}
	return toSerialize, nil
}

type NullableEligibilityPaymentMethodsEntityEntityDetails struct {
	value *EligibilityPaymentMethodsEntityEntityDetails
	isSet bool
}

func (v NullableEligibilityPaymentMethodsEntityEntityDetails) Get() *EligibilityPaymentMethodsEntityEntityDetails {
	return v.value
}

func (v *NullableEligibilityPaymentMethodsEntityEntityDetails) Set(val *EligibilityPaymentMethodsEntityEntityDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityPaymentMethodsEntityEntityDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityPaymentMethodsEntityEntityDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityPaymentMethodsEntityEntityDetails(val *EligibilityPaymentMethodsEntityEntityDetails) *NullableEligibilityPaymentMethodsEntityEntityDetails {
	return &NullableEligibilityPaymentMethodsEntityEntityDetails{value: val, isSet: true}
}

func (v NullableEligibilityPaymentMethodsEntityEntityDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEligibilityPaymentMethodsEntityEntityDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

