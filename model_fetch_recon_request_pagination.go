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

// checks if the FetchReconRequestPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchReconRequestPagination{}

// FetchReconRequestPagination To fetch the next set of settlements, pass the cursor received in the response to the next API call.   To receive the data for the first time, pass the cursor as null.   Limit would be number of settlements that you want to receive.
type FetchReconRequestPagination struct {
	// Number of settlements you want to fetch in the next iteration. Maximum limit is 1000, default value is 10.
	Limit int32 `json:"limit"`
	// Specifies from where the next set of settlement details should be fetched.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewFetchReconRequestPagination instantiates a new FetchReconRequestPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchReconRequestPagination(limit int32) *FetchReconRequestPagination {
	this := FetchReconRequestPagination{}
	this.Limit = limit
	return &this
}

// NewFetchReconRequestPaginationWithDefaults instantiates a new FetchReconRequestPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchReconRequestPaginationWithDefaults() *FetchReconRequestPagination {
	this := FetchReconRequestPagination{}
	return &this
}

// GetLimit returns the Limit field value
func (o *FetchReconRequestPagination) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *FetchReconRequestPagination) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *FetchReconRequestPagination) SetLimit(v int32) {
	o.Limit = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchReconRequestPagination) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchReconRequestPagination) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *FetchReconRequestPagination) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *FetchReconRequestPagination) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *FetchReconRequestPagination) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *FetchReconRequestPagination) UnsetCursor() {
	o.Cursor.Unset()
}

func (o FetchReconRequestPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchReconRequestPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return toSerialize, nil
}

type NullableFetchReconRequestPagination struct {
	value *FetchReconRequestPagination
	isSet bool
}

func (v NullableFetchReconRequestPagination) Get() *FetchReconRequestPagination {
	return v.value
}

func (v *NullableFetchReconRequestPagination) Set(val *FetchReconRequestPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchReconRequestPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchReconRequestPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchReconRequestPagination(val *FetchReconRequestPagination) *NullableFetchReconRequestPagination {
	return &NullableFetchReconRequestPagination{value: val, isSet: true}
}

func (v NullableFetchReconRequestPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchReconRequestPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

