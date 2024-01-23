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

// checks if the ReconEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReconEntity{}

// ReconEntity Settlement detailed recon response
type ReconEntity struct {
	// Specifies from where the next set of settlement details should be fetched.
	Cursor *string `json:"cursor,omitempty"`
	// Number of settlements you want to fetch in the next iteration.
	Limit *int32 `json:"limit,omitempty"`
	Data []ReconEntityDataInner `json:"data,omitempty"`
}


func (o ReconEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReconEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}



