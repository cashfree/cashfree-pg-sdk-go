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

// checks if the InstrumentWebhookDataEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstrumentWebhookDataEntity{}

// InstrumentWebhookDataEntity data entity in webhook
type InstrumentWebhookDataEntity struct {
	Instrument *InstrumentEntity `json:"instrument,omitempty"`
}


func (o InstrumentWebhookDataEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstrumentWebhookDataEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instrument) {
		toSerialize["instrument"] = o.Instrument
	}
	return toSerialize, nil
}


