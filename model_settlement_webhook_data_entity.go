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

// checks if the SettlementWebhookDataEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementWebhookDataEntity{}

// SettlementWebhookDataEntity data entity in webhook
type SettlementWebhookDataEntity struct {
	Settlement *SettlementEntity `json:"settlement,omitempty"`
}


func (o SettlementWebhookDataEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementWebhookDataEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settlement) {
		toSerialize["settlement"] = o.Settlement
	}
	return toSerialize, nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}