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

// checks if the PaymentMethodsFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsFilters{}

// PaymentMethodsFilters Filter for specific Payment Methods
type PaymentMethodsFilters struct {
	// Array of payment methods to be filtered. This is optional, by default all payment methods will be returned. Possible values in [ 'debit_card', 'credit_card', 'prepaid_card', 'corporate_credit_card', 'upi', 'wallet', 'netbanking', 'banktransfer', 'paylater', 'paypal', 'debit_card_emi', 'credit_card_emi', 'upi_credit_card', 'upi_ppi', 'cardless_emi', 'account_based_payment' ] 
	PaymentMethods []string `json:"payment_methods,omitempty"`
}

// NewPaymentMethodsFilters instantiates a new PaymentMethodsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsFilters() *PaymentMethodsFilters {
	this := PaymentMethodsFilters{}
	return &this
}

// NewPaymentMethodsFiltersWithDefaults instantiates a new PaymentMethodsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsFiltersWithDefaults() *PaymentMethodsFilters {
	this := PaymentMethodsFilters{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsFilters) GetPaymentMethods() []string {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret []string
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsFilters) GetPaymentMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsFilters) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []string and assigns it to the PaymentMethods field.
func (o *PaymentMethodsFilters) SetPaymentMethods(v []string) {
	o.PaymentMethods = v
}

func (o PaymentMethodsFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return toSerialize, nil
}

type NullablePaymentMethodsFilters struct {
	value *PaymentMethodsFilters
	isSet bool
}

func (v NullablePaymentMethodsFilters) Get() *PaymentMethodsFilters {
	return v.value
}

func (v *NullablePaymentMethodsFilters) Set(val *PaymentMethodsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsFilters(val *PaymentMethodsFilters) *NullablePaymentMethodsFilters {
	return &NullablePaymentMethodsFilters{value: val, isSet: true}
}

func (v NullablePaymentMethodsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


