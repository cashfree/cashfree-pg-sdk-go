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

// checks if the EligibilityFetchOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityFetchOffersRequest{}

// EligibilityFetchOffersRequest Eligiblty API request
type EligibilityFetchOffersRequest struct {
	Queries OfferQueries `json:"queries"`
	Filters *OfferFilters `json:"filters,omitempty"`
}

// NewEligibilityFetchOffersRequest instantiates a new EligibilityFetchOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityFetchOffersRequest(queries OfferQueries) *EligibilityFetchOffersRequest {
	this := EligibilityFetchOffersRequest{}
	this.Queries = queries
	return &this
}

// NewEligibilityFetchOffersRequestWithDefaults instantiates a new EligibilityFetchOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityFetchOffersRequestWithDefaults() *EligibilityFetchOffersRequest {
	this := EligibilityFetchOffersRequest{}
	return &this
}

// GetQueries returns the Queries field value
func (o *EligibilityFetchOffersRequest) GetQueries() OfferQueries {
	if o == nil {
		var ret OfferQueries
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *EligibilityFetchOffersRequest) GetQueriesOk() (*OfferQueries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value
func (o *EligibilityFetchOffersRequest) SetQueries(v OfferQueries) {
	o.Queries = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *EligibilityFetchOffersRequest) GetFilters() OfferFilters {
	if o == nil || IsNil(o.Filters) {
		var ret OfferFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityFetchOffersRequest) GetFiltersOk() (*OfferFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *EligibilityFetchOffersRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given OfferFilters and assigns it to the Filters field.
func (o *EligibilityFetchOffersRequest) SetFilters(v OfferFilters) {
	o.Filters = &v
}

func (o EligibilityFetchOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityFetchOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queries"] = o.Queries
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableEligibilityFetchOffersRequest struct {
	value *EligibilityFetchOffersRequest
	isSet bool
}

func (v NullableEligibilityFetchOffersRequest) Get() *EligibilityFetchOffersRequest {
	return v.value
}

func (v *NullableEligibilityFetchOffersRequest) Set(val *EligibilityFetchOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityFetchOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityFetchOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityFetchOffersRequest(val *EligibilityFetchOffersRequest) *NullableEligibilityFetchOffersRequest {
	return &NullableEligibilityFetchOffersRequest{value: val, isSet: true}
}

func (v NullableEligibilityFetchOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEligibilityFetchOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


