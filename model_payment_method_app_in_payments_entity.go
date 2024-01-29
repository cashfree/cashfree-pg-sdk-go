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

// checks if the PaymentMethodAppInPaymentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodAppInPaymentsEntity{}

// PaymentMethodAppInPaymentsEntity payment method app object in payment entity
type PaymentMethodAppInPaymentsEntity struct {
	App *PaymentMethodAppInPaymentsEntityApp `json:"app,omitempty"`
}


func (o PaymentMethodAppInPaymentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodAppInPaymentsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}



