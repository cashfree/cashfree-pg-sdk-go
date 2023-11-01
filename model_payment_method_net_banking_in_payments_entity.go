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

// checks if the PaymentMethodNetBankingInPaymentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodNetBankingInPaymentsEntity{}

// PaymentMethodNetBankingInPaymentsEntity netbanking payment method object for pay
type PaymentMethodNetBankingInPaymentsEntity struct {
	Channel string `json:"channel"`
	NetbankingBankCode int32 `json:"netbanking_bank_code"`
	NetbankingBankName string `json:"netbanking_bank_name"`
}

// NewPaymentMethodNetBankingInPaymentsEntity instantiates a new PaymentMethodNetBankingInPaymentsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodNetBankingInPaymentsEntity(channel string, netbankingBankCode int32, netbankingBankName string) *PaymentMethodNetBankingInPaymentsEntity {
	this := PaymentMethodNetBankingInPaymentsEntity{}
	this.Channel = channel
	this.NetbankingBankCode = netbankingBankCode
	this.NetbankingBankName = netbankingBankName
	return &this
}

// NewPaymentMethodNetBankingInPaymentsEntityWithDefaults instantiates a new PaymentMethodNetBankingInPaymentsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodNetBankingInPaymentsEntityWithDefaults() *PaymentMethodNetBankingInPaymentsEntity {
	this := PaymentMethodNetBankingInPaymentsEntity{}
	return &this
}

// GetChannel returns the Channel field value
func (o *PaymentMethodNetBankingInPaymentsEntity) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodNetBankingInPaymentsEntity) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *PaymentMethodNetBankingInPaymentsEntity) SetChannel(v string) {
	o.Channel = v
}

// GetNetbankingBankCode returns the NetbankingBankCode field value
func (o *PaymentMethodNetBankingInPaymentsEntity) GetNetbankingBankCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetbankingBankCode
}

// GetNetbankingBankCodeOk returns a tuple with the NetbankingBankCode field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodNetBankingInPaymentsEntity) GetNetbankingBankCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetbankingBankCode, true
}

// SetNetbankingBankCode sets field value
func (o *PaymentMethodNetBankingInPaymentsEntity) SetNetbankingBankCode(v int32) {
	o.NetbankingBankCode = v
}

// GetNetbankingBankName returns the NetbankingBankName field value
func (o *PaymentMethodNetBankingInPaymentsEntity) GetNetbankingBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetbankingBankName
}

// GetNetbankingBankNameOk returns a tuple with the NetbankingBankName field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodNetBankingInPaymentsEntity) GetNetbankingBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetbankingBankName, true
}

// SetNetbankingBankName sets field value
func (o *PaymentMethodNetBankingInPaymentsEntity) SetNetbankingBankName(v string) {
	o.NetbankingBankName = v
}

func (o PaymentMethodNetBankingInPaymentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodNetBankingInPaymentsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["netbanking_bank_code"] = o.NetbankingBankCode
	toSerialize["netbanking_bank_name"] = o.NetbankingBankName
	return toSerialize, nil
}

type NullablePaymentMethodNetBankingInPaymentsEntity struct {
	value *PaymentMethodNetBankingInPaymentsEntity
	isSet bool
}

func (v NullablePaymentMethodNetBankingInPaymentsEntity) Get() *PaymentMethodNetBankingInPaymentsEntity {
	return v.value
}

func (v *NullablePaymentMethodNetBankingInPaymentsEntity) Set(val *PaymentMethodNetBankingInPaymentsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodNetBankingInPaymentsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodNetBankingInPaymentsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodNetBankingInPaymentsEntity(val *PaymentMethodNetBankingInPaymentsEntity) *NullablePaymentMethodNetBankingInPaymentsEntity {
	return &NullablePaymentMethodNetBankingInPaymentsEntity{value: val, isSet: true}
}

func (v NullablePaymentMethodNetBankingInPaymentsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodNetBankingInPaymentsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

