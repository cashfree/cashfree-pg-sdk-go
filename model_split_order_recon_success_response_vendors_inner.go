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
	"time"
)

// checks if the SplitOrderReconSuccessResponseVendorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitOrderReconSuccessResponseVendorsInner{}

// SplitOrderReconSuccessResponseVendorsInner struct for SplitOrderReconSuccessResponseVendorsInner
type SplitOrderReconSuccessResponseVendorsInner struct {
	// Unique identifier for the vendor.
	VendorId *string `json:"vendor_id,omitempty"`
	// Settlement ID associated with the vendor.
	SettlementId *int64 `json:"settlement_id,omitempty"`
	// Settlement amount allocated to the vendor.
	SettlementAmount *float32 `json:"settlement_amount,omitempty"`
	// Date and time when the vendor is eligible for the settlement.
	SettlementEligibilityDate *time.Time `json:"settlement_eligibility_date,omitempty"`
}


func (o SplitOrderReconSuccessResponseVendorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitOrderReconSuccessResponseVendorsInner) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.SettlementId) {
		toSerialize["settlement_id"] = o.SettlementId
	}
	if !IsNil(o.SettlementAmount) {
		toSerialize["settlement_amount"] = o.SettlementAmount
	}
	if !IsNil(o.SettlementEligibilityDate) {
		toSerialize["settlement_eligibility_date"] = o.SettlementEligibilityDate
	}
	return toSerialize, nil
}



