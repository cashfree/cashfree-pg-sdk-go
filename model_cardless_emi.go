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

// checks if the CardlessEMI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardlessEMI{}

// CardlessEMI Request body for cardless emi payment method
type CardlessEMI struct {
	// The channel for cardless EMI is always `link`
	Channel *string `json:"channel,omitempty"`
	// One of [`flexmoney`, `zestmoney`, `hdfc`, `icici`, `cashe`, `idfc`, `kotak`]
	Provider *string `json:"provider,omitempty"`
	// Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as 'invalid_request_error' and code as 'invalid_request_error'
	Phone *string `json:"phone,omitempty"`
	// EMI tenure for the selected provider. This is mandatory when provider is one of [`hdfc`, `icici`, `cashe`, `idfc`, `kotak`]
	EmiTenure *int32 `json:"emi_tenure,omitempty"`
}


func (o CardlessEMI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardlessEMI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.EmiTenure) {
		toSerialize["emi_tenure"] = o.EmiTenure
	}
	return toSerialize, nil
}



