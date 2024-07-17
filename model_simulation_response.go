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

// checks if the SimulationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimulationResponse{}

// SimulationResponse Simulation response object
type SimulationResponse struct {
	SimulationId *string `json:"simulation_id,omitempty"`
	Entity *string `json:"entity,omitempty"`
	EntityId *string `json:"entity_id,omitempty"`
	EntitySimulation *EntitySimulationResponse `json:"entity_simulation,omitempty"`
}


func (o SimulationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimulationResponse) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SimulationId) {
		toSerialize["simulation_id"] = o.SimulationId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.EntitySimulation) {
		toSerialize["entity_simulation"] = o.EntitySimulation
	}
	return toSerialize, nil
}


