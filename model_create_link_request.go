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
)

// checks if the CreateLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkRequest{}

// CreateLinkRequest Request paramenters for link creation
type CreateLinkRequest struct {
	// Unique Identifier (provided by merchant) for the Link. Alphanumeric and only - and _ allowed (50 character limit). Use this for other link-related APIs.
	LinkId string `json:"link_id"`
	// Amount to be collected using this link. Provide upto two decimals for paise.
	LinkAmount float64 `json:"link_amount"`
	// Currency for the payment link. Default is INR. Contact care@cashfree.com to enable new currencies.
	LinkCurrency string `json:"link_currency"`
	// A brief description for which payment must be collected. This is shown to the customer.
	LinkPurpose string `json:"link_purpose"`
	CustomerDetails LinkCustomerDetailsEntity `json:"customer_details"`
	// If \"true\", customer can make partial payments for the link.
	LinkPartialPayments *bool `json:"link_partial_payments,omitempty"`
	// Minimum amount in first installment that needs to be paid by the customer if partial payments are enabled. This should be less than the link_amount.
	LinkMinimumPartialAmount *float64 `json:"link_minimum_partial_amount,omitempty"`
	// Time after which the link expires. Customers will not be able to make the payment beyond the time specified here. You can provide them in a valid ISO 8601 time format. Default is 30 days.
	LinkExpiryTime *string `json:"link_expiry_time,omitempty"`
	LinkNotify *LinkNotifyEntity `json:"link_notify,omitempty"`
	// If \"true\", reminders will be sent to customers for collecting payments.
	LinkAutoReminders *bool `json:"link_auto_reminders,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	LinkNotes *map[string]string `json:"link_notes,omitempty"`
	LinkMeta *LinkMetaResponseEntity `json:"link_meta,omitempty"`
}


func (o CreateLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link_id"] = o.LinkId
	toSerialize["link_amount"] = o.LinkAmount
	toSerialize["link_currency"] = o.LinkCurrency
	toSerialize["link_purpose"] = o.LinkPurpose
	toSerialize["customer_details"] = o.CustomerDetails
	if !IsNil(o.LinkPartialPayments) {
		toSerialize["link_partial_payments"] = o.LinkPartialPayments
	}
	if !IsNil(o.LinkMinimumPartialAmount) {
		toSerialize["link_minimum_partial_amount"] = o.LinkMinimumPartialAmount
	}
	if !IsNil(o.LinkExpiryTime) {
		toSerialize["link_expiry_time"] = o.LinkExpiryTime
	}
	if !IsNil(o.LinkNotify) {
		toSerialize["link_notify"] = o.LinkNotify
	}
	if !IsNil(o.LinkAutoReminders) {
		toSerialize["link_auto_reminders"] = o.LinkAutoReminders
	}
	if !IsNil(o.LinkNotes) {
		toSerialize["link_notes"] = o.LinkNotes
	}
	if !IsNil(o.LinkMeta) {
		toSerialize["link_meta"] = o.LinkMeta
	}
	return toSerialize, nil
}



