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

// checks if the PaymentMethodsQueries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsQueries{}

// PaymentMethodsQueries Payment Method Query Object
type PaymentMethodsQueries struct {
	// Amount of the order.
	Amount *float32 `json:"amount,omitempty"`
	// OrderId of the order. Either of `order_id` or `order_amount` is mandatory.
	OrderId *string `json:"order_id,omitempty"`
}


func (o PaymentMethodsQueries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsQueries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	return toSerialize, nil
}



