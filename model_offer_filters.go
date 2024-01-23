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

// checks if the OfferFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferFilters{}

// OfferFilters Filter for offers
type OfferFilters struct {
	// Array of offer_type to be filtered.
	OfferType []OfferType `json:"offer_type,omitempty"`
}


func (o OfferFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OfferType) {
		toSerialize["offer_type"] = o.OfferType
	}
	return toSerialize, nil
}



