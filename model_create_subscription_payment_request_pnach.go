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

// checks if the CreateSubscriptionPaymentRequestPnach type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionPaymentRequestPnach{}

// CreateSubscriptionPaymentRequestPnach payment method pnach.
type CreateSubscriptionPaymentRequestPnach struct {
	// Channel. can be post
	Channel *string `json:"channel,omitempty"`
	// Account holder name
	AccountHolderName *string `json:"account_holder_name,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Account bank code
	AccountBankCode *string `json:"account_bank_code,omitempty"`
	// Account type
	AccountType *string `json:"account_type,omitempty"`
	// Account IFSC
	AccountIfsc *string `json:"account_ifsc,omitempty"`
	// Mandate creation date
	MandateCreationDate *string `json:"mandate_creation_date,omitempty"`
	// Mandate start date
	MandateStartDate *string `json:"mandate_start_date,omitempty"`
}


func (o CreateSubscriptionPaymentRequestPnach) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionPaymentRequestPnach) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.AccountHolderName) {
		toSerialize["account_holder_name"] = o.AccountHolderName
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.AccountBankCode) {
		toSerialize["account_bank_code"] = o.AccountBankCode
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.AccountIfsc) {
		toSerialize["account_ifsc"] = o.AccountIfsc
	}
	if !IsNil(o.MandateCreationDate) {
		toSerialize["mandate_creation_date"] = o.MandateCreationDate
	}
	if !IsNil(o.MandateStartDate) {
		toSerialize["mandate_start_date"] = o.MandateStartDate
	}
	return toSerialize, nil
}


