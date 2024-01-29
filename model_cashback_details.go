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

// checks if the CashbackDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashbackDetails{}

// CashbackDetails Cashback detail boject
type CashbackDetails struct {
	// Type of discount
	CashbackType string `json:"cashback_type"`
	// Value of Discount.
	CashbackValue float32 `json:"cashback_value"`
	// Maximum Value of Cashback allowed.
	MaxCashbackAmount float32 `json:"max_cashback_amount"`
}


func (o CashbackDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashbackDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["cashback_type"] = o.CashbackType
	toSerialize["cashback_value"] = o.CashbackValue
	toSerialize["max_cashback_amount"] = o.MaxCashbackAmount
	return toSerialize, nil
}



