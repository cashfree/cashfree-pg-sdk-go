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

// checks if the OfferDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferDetails{}

// OfferDetails Offer details and type
type OfferDetails struct {
	// Offer Type for the Offer.
	OfferType string `json:"offer_type"`
	DiscountDetails *DiscountDetails `json:"discount_details,omitempty"`
	CashbackDetails *CashbackDetails `json:"cashback_details,omitempty"`
}

// NewOfferDetails instantiates a new OfferDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferDetails(offerType string) *OfferDetails {
	this := OfferDetails{}
	this.OfferType = offerType
	return &this
}

// NewOfferDetailsWithDefaults instantiates a new OfferDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferDetailsWithDefaults() *OfferDetails {
	this := OfferDetails{}
	return &this
}

// GetOfferType returns the OfferType field value
func (o *OfferDetails) GetOfferType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferType
}

// GetOfferTypeOk returns a tuple with the OfferType field value
// and a boolean to check if the value has been set.
func (o *OfferDetails) GetOfferTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferType, true
}

// SetOfferType sets field value
func (o *OfferDetails) SetOfferType(v string) {
	o.OfferType = v
}

// GetDiscountDetails returns the DiscountDetails field value if set, zero value otherwise.
func (o *OfferDetails) GetDiscountDetails() DiscountDetails {
	if o == nil || IsNil(o.DiscountDetails) {
		var ret DiscountDetails
		return ret
	}
	return *o.DiscountDetails
}

// GetDiscountDetailsOk returns a tuple with the DiscountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferDetails) GetDiscountDetailsOk() (*DiscountDetails, bool) {
	if o == nil || IsNil(o.DiscountDetails) {
		return nil, false
	}
	return o.DiscountDetails, true
}

// HasDiscountDetails returns a boolean if a field has been set.
func (o *OfferDetails) HasDiscountDetails() bool {
	if o != nil && !IsNil(o.DiscountDetails) {
		return true
	}

	return false
}

// SetDiscountDetails gets a reference to the given DiscountDetails and assigns it to the DiscountDetails field.
func (o *OfferDetails) SetDiscountDetails(v DiscountDetails) {
	o.DiscountDetails = &v
}

// GetCashbackDetails returns the CashbackDetails field value if set, zero value otherwise.
func (o *OfferDetails) GetCashbackDetails() CashbackDetails {
	if o == nil || IsNil(o.CashbackDetails) {
		var ret CashbackDetails
		return ret
	}
	return *o.CashbackDetails
}

// GetCashbackDetailsOk returns a tuple with the CashbackDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferDetails) GetCashbackDetailsOk() (*CashbackDetails, bool) {
	if o == nil || IsNil(o.CashbackDetails) {
		return nil, false
	}
	return o.CashbackDetails, true
}

// HasCashbackDetails returns a boolean if a field has been set.
func (o *OfferDetails) HasCashbackDetails() bool {
	if o != nil && !IsNil(o.CashbackDetails) {
		return true
	}

	return false
}

// SetCashbackDetails gets a reference to the given CashbackDetails and assigns it to the CashbackDetails field.
func (o *OfferDetails) SetCashbackDetails(v CashbackDetails) {
	o.CashbackDetails = &v
}

func (o OfferDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offer_type"] = o.OfferType
	if !IsNil(o.DiscountDetails) {
		toSerialize["discount_details"] = o.DiscountDetails
	}
	if !IsNil(o.CashbackDetails) {
		toSerialize["cashback_details"] = o.CashbackDetails
	}
	return toSerialize, nil
}

type NullableOfferDetails struct {
	value *OfferDetails
	isSet bool
}

func (v NullableOfferDetails) Get() *OfferDetails {
	return v.value
}

func (v *NullableOfferDetails) Set(val *OfferDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferDetails(val *OfferDetails) *NullableOfferDetails {
	return &NullableOfferDetails{value: val, isSet: true}
}

func (v NullableOfferDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

