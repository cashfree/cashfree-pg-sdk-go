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

// checks if the OrderEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderEntity{}

// OrderEntity The complete order entity
type OrderEntity struct {
	// unique id generated by cashfree for your order
	CfOrderId *string `json:"cf_order_id,omitempty"`
	// order_id sent during the api request
	OrderId *string `json:"order_id,omitempty"`
	// Type of the entity.
	Entity *string `json:"entity,omitempty"`
	// Currency of the order. Example INR
	OrderCurrency *string `json:"order_currency,omitempty"`
	OrderAmount *float32 `json:"order_amount,omitempty"`
	// Possible values are  - `ACTIVE`: Order does not have a sucessful transaction yet - `PAID`: Order is PAID with one successful transaction - `EXPIRED`: Order was not PAID and not it has expired. No transaction can be initiated for an EXPIRED order. `TERMINATED`: Order terminated `TERMINATION_REQUESTED`: Order termination requested
	OrderStatus *string `json:"order_status,omitempty"`
	PaymentSessionId *string `json:"payment_session_id,omitempty"`
	OrderExpiryTime *time.Time `json:"order_expiry_time,omitempty"`
	// Additional note for order
	OrderNote *string `json:"order_note,omitempty"`
	// When the order was created at cashfree's server
	CreatedAt *time.Time `json:"created_at,omitempty"`
	OrderSplits []VendorSplit `json:"order_splits,omitempty"`
	CustomerDetails *CustomerDetails `json:"customer_details,omitempty"`
	OrderMeta *OrderMeta `json:"order_meta,omitempty"`
	// Custom Tags in thr form of {\"key\":\"value\"} which can be passed for an order. A maximum of 10 tags can be added
	OrderTags *map[string]string `json:"order_tags,omitempty"`
}


func (o OrderEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfOrderId) {
		toSerialize["cf_order_id"] = o.CfOrderId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.OrderCurrency) {
		toSerialize["order_currency"] = o.OrderCurrency
	}
	if !IsNil(o.OrderAmount) {
		toSerialize["order_amount"] = o.OrderAmount
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.PaymentSessionId) {
		toSerialize["payment_session_id"] = o.PaymentSessionId
	}
	if !IsNil(o.OrderExpiryTime) {
		toSerialize["order_expiry_time"] = o.OrderExpiryTime
	}
	if !IsNil(o.OrderNote) {
		toSerialize["order_note"] = o.OrderNote
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.OrderSplits) {
		toSerialize["order_splits"] = o.OrderSplits
	}
	if !IsNil(o.CustomerDetails) {
		toSerialize["customer_details"] = o.CustomerDetails
	}
	if !IsNil(o.OrderMeta) {
		toSerialize["order_meta"] = o.OrderMeta
	}
	if !IsNil(o.OrderTags) {
		toSerialize["order_tags"] = o.OrderTags
	}
	return toSerialize, nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}