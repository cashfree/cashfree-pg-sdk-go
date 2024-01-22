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

// checks if the LinkEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkEntity{}

// LinkEntity Payment link success creation response object
type LinkEntity struct {
	CfLinkId *string `json:"cf_link_id,omitempty"`
	LinkId *string `json:"link_id,omitempty"`
	LinkStatus *string `json:"link_status,omitempty"`
	LinkCurrency *string `json:"link_currency,omitempty"`
	LinkAmount *float32 `json:"link_amount,omitempty"`
	LinkAmountPaid *float32 `json:"link_amount_paid,omitempty"`
	LinkPartialPayments *bool `json:"link_partial_payments,omitempty"`
	LinkMinimumPartialAmount *float32 `json:"link_minimum_partial_amount,omitempty"`
	LinkPurpose *string `json:"link_purpose,omitempty"`
	LinkCreatedAt *string `json:"link_created_at,omitempty"`
	CustomerDetails *LinkCustomerDetailsEntity `json:"customer_details,omitempty"`
	LinkMeta *LinkMetaResponseEntity `json:"link_meta,omitempty"`
	LinkUrl *string `json:"link_url,omitempty"`
	LinkExpiryTime *string `json:"link_expiry_time,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	LinkNotes *map[string]string `json:"link_notes,omitempty"`
	LinkAutoReminders *bool `json:"link_auto_reminders,omitempty"`
	LinkNotify *LinkNotifyEntity `json:"link_notify,omitempty"`
}


func (o LinkEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfLinkId) {
		toSerialize["cf_link_id"] = o.CfLinkId
	}
	if !IsNil(o.LinkId) {
		toSerialize["link_id"] = o.LinkId
	}
	if !IsNil(o.LinkStatus) {
		toSerialize["link_status"] = o.LinkStatus
	}
	if !IsNil(o.LinkCurrency) {
		toSerialize["link_currency"] = o.LinkCurrency
	}
	if !IsNil(o.LinkAmount) {
		toSerialize["link_amount"] = o.LinkAmount
	}
	if !IsNil(o.LinkAmountPaid) {
		toSerialize["link_amount_paid"] = o.LinkAmountPaid
	}
	if !IsNil(o.LinkPartialPayments) {
		toSerialize["link_partial_payments"] = o.LinkPartialPayments
	}
	if !IsNil(o.LinkMinimumPartialAmount) {
		toSerialize["link_minimum_partial_amount"] = o.LinkMinimumPartialAmount
	}
	if !IsNil(o.LinkPurpose) {
		toSerialize["link_purpose"] = o.LinkPurpose
	}
	if !IsNil(o.LinkCreatedAt) {
		toSerialize["link_created_at"] = o.LinkCreatedAt
	}
	if !IsNil(o.CustomerDetails) {
		toSerialize["customer_details"] = o.CustomerDetails
	}
	if !IsNil(o.LinkMeta) {
		toSerialize["link_meta"] = o.LinkMeta
	}
	if !IsNil(o.LinkUrl) {
		toSerialize["link_url"] = o.LinkUrl
	}
	if !IsNil(o.LinkExpiryTime) {
		toSerialize["link_expiry_time"] = o.LinkExpiryTime
	}
	if !IsNil(o.LinkNotes) {
		toSerialize["link_notes"] = o.LinkNotes
	}
	if !IsNil(o.LinkAutoReminders) {
		toSerialize["link_auto_reminders"] = o.LinkAutoReminders
	}
	if !IsNil(o.LinkNotify) {
		toSerialize["link_notify"] = o.LinkNotify
	}
	return toSerialize, nil
}



