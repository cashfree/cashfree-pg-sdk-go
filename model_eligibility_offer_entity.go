/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the EligibilityOfferEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityOfferEntity{}

// EligibilityOfferEntity Eligible offer object
type EligibilityOfferEntity struct {
	Eligibility *bool `json:"eligibility,omitempty"`
	EntityType *string `json:"entity_type,omitempty"`
	EntityValue *string `json:"entity_value,omitempty"`
	EntityDetails *OfferEntity `json:"entity_details,omitempty"`
}


func (o EligibilityOfferEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityOfferEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Eligibility) {
		toSerialize["eligibility"] = o.Eligibility
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.EntityValue) {
		toSerialize["entity_value"] = o.EntityValue
	}
	if !IsNil(o.EntityDetails) {
		toSerialize["entity_details"] = o.EntityDetails
	}
	return toSerialize, nil
}



