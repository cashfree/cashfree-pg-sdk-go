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

// checks if the ApiError404 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiError404{}

// ApiError404 Error when resource requested is not found
type ApiError404 struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
	// invalid_request_error
	Type *string `json:"type,omitempty"`
}

// NewApiError404 instantiates a new ApiError404 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError404() *ApiError404 {
	this := ApiError404{}
	return &this
}

// NewApiError404WithDefaults instantiates a new ApiError404 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiError404WithDefaults() *ApiError404 {
	this := ApiError404{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiError404) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError404) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiError404) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiError404) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiError404) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError404) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiError404) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiError404) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiError404) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError404) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiError404) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiError404) SetType(v string) {
	o.Type = &v
}

func (o ApiError404) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiError404) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApiError404 struct {
	value *ApiError404
	isSet bool
}

func (v NullableApiError404) Get() *ApiError404 {
	return v.value
}

func (v *NullableApiError404) Set(val *ApiError404) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError404) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError404) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError404(val *ApiError404) *NullableApiError404 {
	return &NullableApiError404{value: val, isSet: true}
}

func (v NullableApiError404) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError404) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


