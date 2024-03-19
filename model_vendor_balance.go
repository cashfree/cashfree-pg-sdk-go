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

// checks if the VendorBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorBalance{}

// VendorBalance Vendor Balance entity object
type VendorBalance struct {
	MerchantId *float32 `json:"merchant_id,omitempty"`
	VendorId *string `json:"vendor_id,omitempty"`
	MerchantUnsettled *float32 `json:"merchant_unsettled,omitempty"`
	VendorUnsettled *float32 `json:"vendor_unsettled,omitempty"`
}


func (o VendorBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorBalance) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.MerchantUnsettled) {
		toSerialize["merchant_unsettled"] = o.MerchantUnsettled
	}
	if !IsNil(o.VendorUnsettled) {
		toSerialize["vendor_unsettled"] = o.VendorUnsettled
	}
	return toSerialize, nil
}



