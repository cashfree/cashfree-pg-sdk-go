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

// checks if the CardlessEMIEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardlessEMIEntity{}

// CardlessEMIEntity cardless EMI object
type CardlessEMIEntity struct {
	PaymentMethod *string `json:"payment_method,omitempty"`
	EmiPlans []EMIPlansArray `json:"emi_plans,omitempty"`
}

// NewCardlessEMIEntity instantiates a new CardlessEMIEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardlessEMIEntity() *CardlessEMIEntity {
	this := CardlessEMIEntity{}
	return &this
}

// NewCardlessEMIEntityWithDefaults instantiates a new CardlessEMIEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardlessEMIEntityWithDefaults() *CardlessEMIEntity {
	this := CardlessEMIEntity{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CardlessEMIEntity) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardlessEMIEntity) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CardlessEMIEntity) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *CardlessEMIEntity) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetEmiPlans returns the EmiPlans field value if set, zero value otherwise.
func (o *CardlessEMIEntity) GetEmiPlans() []EMIPlansArray {
	if o == nil || IsNil(o.EmiPlans) {
		var ret []EMIPlansArray
		return ret
	}
	return o.EmiPlans
}

// GetEmiPlansOk returns a tuple with the EmiPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardlessEMIEntity) GetEmiPlansOk() ([]EMIPlansArray, bool) {
	if o == nil || IsNil(o.EmiPlans) {
		return nil, false
	}
	return o.EmiPlans, true
}

// HasEmiPlans returns a boolean if a field has been set.
func (o *CardlessEMIEntity) HasEmiPlans() bool {
	if o != nil && !IsNil(o.EmiPlans) {
		return true
	}

	return false
}

// SetEmiPlans gets a reference to the given []EMIPlansArray and assigns it to the EmiPlans field.
func (o *CardlessEMIEntity) SetEmiPlans(v []EMIPlansArray) {
	o.EmiPlans = v
}

func (o CardlessEMIEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardlessEMIEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.EmiPlans) {
		toSerialize["emi_plans"] = o.EmiPlans
	}
	return toSerialize, nil
}

type NullableCardlessEMIEntity struct {
	value *CardlessEMIEntity
	isSet bool
}

func (v NullableCardlessEMIEntity) Get() *CardlessEMIEntity {
	return v.value
}

func (v *NullableCardlessEMIEntity) Set(val *CardlessEMIEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCardlessEMIEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCardlessEMIEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardlessEMIEntity(val *CardlessEMIEntity) *NullableCardlessEMIEntity {
	return &NullableCardlessEMIEntity{value: val, isSet: true}
}

func (v NullableCardlessEMIEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardlessEMIEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


