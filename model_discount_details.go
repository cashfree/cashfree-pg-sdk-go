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

// checks if the DiscountDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountDetails{}

// DiscountDetails detils of the discount object of offer
type DiscountDetails struct {
	// Type of discount
	DiscountType string `json:"discount_type"`
	// Value of Discount.
	DiscountValue string `json:"discount_value"`
	// Maximum Value of Discount allowed.
	MaxDiscountAmount string `json:"max_discount_amount"`
}

// NewDiscountDetails instantiates a new DiscountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountDetails(discountType string, discountValue string, maxDiscountAmount string) *DiscountDetails {
	this := DiscountDetails{}
	this.DiscountType = discountType
	this.DiscountValue = discountValue
	this.MaxDiscountAmount = maxDiscountAmount
	return &this
}

// NewDiscountDetailsWithDefaults instantiates a new DiscountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountDetailsWithDefaults() *DiscountDetails {
	this := DiscountDetails{}
	return &this
}

// GetDiscountType returns the DiscountType field value
func (o *DiscountDetails) GetDiscountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetDiscountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountType, true
}

// SetDiscountType sets field value
func (o *DiscountDetails) SetDiscountType(v string) {
	o.DiscountType = v
}

// GetDiscountValue returns the DiscountValue field value
func (o *DiscountDetails) GetDiscountValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountValue
}

// GetDiscountValueOk returns a tuple with the DiscountValue field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetDiscountValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountValue, true
}

// SetDiscountValue sets field value
func (o *DiscountDetails) SetDiscountValue(v string) {
	o.DiscountValue = v
}

// GetMaxDiscountAmount returns the MaxDiscountAmount field value
func (o *DiscountDetails) GetMaxDiscountAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxDiscountAmount
}

// GetMaxDiscountAmountOk returns a tuple with the MaxDiscountAmount field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetMaxDiscountAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxDiscountAmount, true
}

// SetMaxDiscountAmount sets field value
func (o *DiscountDetails) SetMaxDiscountAmount(v string) {
	o.MaxDiscountAmount = v
}

func (o DiscountDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discount_type"] = o.DiscountType
	toSerialize["discount_value"] = o.DiscountValue
	toSerialize["max_discount_amount"] = o.MaxDiscountAmount
	return toSerialize, nil
}

type NullableDiscountDetails struct {
	value *DiscountDetails
	isSet bool
}

func (v NullableDiscountDetails) Get() *DiscountDetails {
	return v.value
}

func (v *NullableDiscountDetails) Set(val *DiscountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountDetails(val *DiscountDetails) *NullableDiscountDetails {
	return &NullableDiscountDetails{value: val, isSet: true}
}

func (v NullableDiscountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


