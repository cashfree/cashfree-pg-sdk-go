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

// checks if the LinkCustomerDetailsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkCustomerDetailsEntity{}

// LinkCustomerDetailsEntity Payment link customer entity
type LinkCustomerDetailsEntity struct {
	// Customer phone number
	CustomerPhone string `json:"customer_phone"`
	// Customer email address
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Customer name
	CustomerName *string `json:"customer_name,omitempty"`
}

// NewLinkCustomerDetailsEntity instantiates a new LinkCustomerDetailsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkCustomerDetailsEntity(customerPhone string) *LinkCustomerDetailsEntity {
	this := LinkCustomerDetailsEntity{}
	this.CustomerPhone = customerPhone
	return &this
}

// NewLinkCustomerDetailsEntityWithDefaults instantiates a new LinkCustomerDetailsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkCustomerDetailsEntityWithDefaults() *LinkCustomerDetailsEntity {
	this := LinkCustomerDetailsEntity{}
	return &this
}

// GetCustomerPhone returns the CustomerPhone field value
func (o *LinkCustomerDetailsEntity) GetCustomerPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerPhone
}

// GetCustomerPhoneOk returns a tuple with the CustomerPhone field value
// and a boolean to check if the value has been set.
func (o *LinkCustomerDetailsEntity) GetCustomerPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerPhone, true
}

// SetCustomerPhone sets field value
func (o *LinkCustomerDetailsEntity) SetCustomerPhone(v string) {
	o.CustomerPhone = v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *LinkCustomerDetailsEntity) GetCustomerEmail() string {
	if o == nil || IsNil(o.CustomerEmail) {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCustomerDetailsEntity) GetCustomerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *LinkCustomerDetailsEntity) HasCustomerEmail() bool {
	if o != nil && !IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *LinkCustomerDetailsEntity) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *LinkCustomerDetailsEntity) GetCustomerName() string {
	if o == nil || IsNil(o.CustomerName) {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCustomerDetailsEntity) GetCustomerNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerName) {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *LinkCustomerDetailsEntity) HasCustomerName() bool {
	if o != nil && !IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *LinkCustomerDetailsEntity) SetCustomerName(v string) {
	o.CustomerName = &v
}

func (o LinkCustomerDetailsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkCustomerDetailsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_phone"] = o.CustomerPhone
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	return toSerialize, nil
}

type NullableLinkCustomerDetailsEntity struct {
	value *LinkCustomerDetailsEntity
	isSet bool
}

func (v NullableLinkCustomerDetailsEntity) Get() *LinkCustomerDetailsEntity {
	return v.value
}

func (v *NullableLinkCustomerDetailsEntity) Set(val *LinkCustomerDetailsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkCustomerDetailsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkCustomerDetailsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkCustomerDetailsEntity(val *LinkCustomerDetailsEntity) *NullableLinkCustomerDetailsEntity {
	return &NullableLinkCustomerDetailsEntity{value: val, isSet: true}
}

func (v NullableLinkCustomerDetailsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkCustomerDetailsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


