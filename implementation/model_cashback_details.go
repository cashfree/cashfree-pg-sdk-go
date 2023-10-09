/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-09-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"encoding/json"
)

// checks if the CashbackDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashbackDetails{}

// CashbackDetails struct for CashbackDetails
type CashbackDetails struct {
	// Type of discount
	CashbackType string `json:"cashback_type,omitempty"`
	// Value of Discount.
	CashbackValue float32 `json:"cashback_value,omitempty"`
	// Maximum Value of Cashback allowed.
	MaxCashbackAmount float32 `json:"max_cashback_amount"`
}

// NewCashbackDetails instantiates a new CashbackDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashbackDetails(maxCashbackAmount float32) *CashbackDetails {
	this := CashbackDetails{}
	this.MaxCashbackAmount = maxCashbackAmount
	return &this
}

// NewCashbackDetailsWithDefaults instantiates a new CashbackDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashbackDetailsWithDefaults() *CashbackDetails {
	this := CashbackDetails{}
	return &this
}

// GetCashbackType returns the CashbackType field value if set, zero value otherwise.
func (o *CashbackDetails) GetCashbackType() string {
	if o == nil || IsNil(o.CashbackType) {
		var ret string
		return ret
	}
	return o.CashbackType
}

// GetCashbackTypeOk returns a tuple with the CashbackType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashbackDetails) GetCashbackTypeOk() (string, bool) {
	if o == nil || IsNil(o.CashbackType) {
		return "", false
	}
	return o.CashbackType, true
}

// HasCashbackType returns a boolean if a field has been set.
func (o *CashbackDetails) HasCashbackType() bool {
	if o != nil && !IsNil(o.CashbackType) {
		return true
	}

	return false
}

// SetCashbackType gets a reference to the given string and assigns it to the CashbackType field.
func (o *CashbackDetails) SetCashbackType(v string) {
	o.CashbackType = v
}

// GetCashbackValue returns the CashbackValue field value if set, zero value otherwise.
func (o *CashbackDetails) GetCashbackValue() float32 {
	if o == nil || IsNil(o.CashbackValue) {
		var ret float32
		return ret
	}
	return o.CashbackValue
}

// GetCashbackValueOk returns a tuple with the CashbackValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashbackDetails) GetCashbackValueOk() (float32, bool) {
	if o == nil || IsNil(o.CashbackValue) {
		return 0.0, false
	}
	return o.CashbackValue, true
}

// HasCashbackValue returns a boolean if a field has been set.
func (o *CashbackDetails) HasCashbackValue() bool {
	if o != nil && !IsNil(o.CashbackValue) {
		return true
	}

	return false
}

// SetCashbackValue gets a reference to the given string and assigns it to the CashbackValue field.
func (o *CashbackDetails) SetCashbackValue(v float32) {
	o.CashbackValue = v
}

// GetMaxCashbackAmount returns the MaxCashbackAmount field value
func (o *CashbackDetails) GetMaxCashbackAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxCashbackAmount
}

// GetMaxCashbackAmountOk returns a tuple with the MaxCashbackAmount field value
// and a boolean to check if the value has been set.
func (o *CashbackDetails) GetMaxCashbackAmountOk() (float32, bool) {
	if o == nil {
		return 0.0, false
	}
	return o.MaxCashbackAmount, true
}

// SetMaxCashbackAmount sets field value
func (o *CashbackDetails) SetMaxCashbackAmount(v float32) {
	o.MaxCashbackAmount = v
}

func (o CashbackDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashbackDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CashbackType) {
		toSerialize["cashback_type"] = o.CashbackType
	}
	if !IsNil(o.CashbackValue) {
		toSerialize["cashback_value"] = o.CashbackValue
	}
	toSerialize["max_cashback_amount"] = o.MaxCashbackAmount
	return toSerialize, nil
}

type NullableCashbackDetails struct {
	value *CashbackDetails
	isSet bool
}

func (v NullableCashbackDetails) Get() *CashbackDetails {
	return v.value
}

func (v *NullableCashbackDetails) Set(val *CashbackDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCashbackDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCashbackDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashbackDetails(val *CashbackDetails) *NullableCashbackDetails {
	return &NullableCashbackDetails{value: val, isSet: true}
}

func (v NullableCashbackDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashbackDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


