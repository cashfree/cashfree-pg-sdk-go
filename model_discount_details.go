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

// checks if the DiscountDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountDetails{}

// DiscountDetails detils of the discount object of offer
type DiscountDetails struct {
	// Type of discount
	DiscountType string `json:"discount_type"`
	// Value of Discount.
	DiscountValue float32 `json:"discount_value"`
	// Maximum Value of Discount allowed.
	MaxDiscountAmount float32 `json:"max_discount_amount"`
}


func (o DiscountDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discount_type"] = o.DiscountType
	toSerialize["discount_value"] = o.DiscountValue
	toSerialize["max_discount_amount"] = o.MaxDiscountAmount
	return toSerialize, nil
}



