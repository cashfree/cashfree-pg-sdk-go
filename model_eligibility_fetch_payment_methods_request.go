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

// checks if the EligibilityFetchPaymentMethodsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityFetchPaymentMethodsRequest{}

// EligibilityFetchPaymentMethodsRequest eligibilty request to find eligible payment method
type EligibilityFetchPaymentMethodsRequest struct {
	Queries PaymentMethodsQueries `json:"queries"`
	Filters *PaymentMethodsFilters `json:"filters,omitempty"`
}


func (o EligibilityFetchPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityFetchPaymentMethodsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queries"] = o.Queries
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}