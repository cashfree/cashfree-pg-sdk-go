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

// checks if the NetBankingPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetBankingPaymentMethod{}

// NetBankingPaymentMethod Payment method for netbanking object
type NetBankingPaymentMethod struct {
	Netbanking Netbanking `json:"netbanking"`
}

// NewNetBankingPaymentMethod instantiates a new NetBankingPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetBankingPaymentMethod(netbanking Netbanking) *NetBankingPaymentMethod {
	this := NetBankingPaymentMethod{}
	this.Netbanking = netbanking
	return &this
}

// NewNetBankingPaymentMethodWithDefaults instantiates a new NetBankingPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetBankingPaymentMethodWithDefaults() *NetBankingPaymentMethod {
	this := NetBankingPaymentMethod{}
	return &this
}

// GetNetbanking returns the Netbanking field value
func (o *NetBankingPaymentMethod) GetNetbanking() Netbanking {
	if o == nil {
		var ret Netbanking
		return ret
	}

	return o.Netbanking
}

// GetNetbankingOk returns a tuple with the Netbanking field value
// and a boolean to check if the value has been set.
func (o *NetBankingPaymentMethod) GetNetbankingOk() (*Netbanking, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netbanking, true
}

// SetNetbanking sets field value
func (o *NetBankingPaymentMethod) SetNetbanking(v Netbanking) {
	o.Netbanking = v
}

func (o NetBankingPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetBankingPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["netbanking"] = o.Netbanking
	return toSerialize, nil
}

type NullableNetBankingPaymentMethod struct {
	value *NetBankingPaymentMethod
	isSet bool
}

func (v NullableNetBankingPaymentMethod) Get() *NetBankingPaymentMethod {
	return v.value
}

func (v *NullableNetBankingPaymentMethod) Set(val *NetBankingPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableNetBankingPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableNetBankingPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetBankingPaymentMethod(val *NetBankingPaymentMethod) *NullableNetBankingPaymentMethod {
	return &NullableNetBankingPaymentMethod{value: val, isSet: true}
}

func (v NullableNetBankingPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetBankingPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


