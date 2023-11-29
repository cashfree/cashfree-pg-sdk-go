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

// checks if the RefundEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundEntity{}

// RefundEntity The refund entity
type RefundEntity struct {
	// Cashfree Payments ID of the payment for which refund is initiated
	CfPaymentId *int32 `json:"cf_payment_id,omitempty"`
	// Cashfree Payments ID for a refund
	CfRefundId *string `json:"cf_refund_id,omitempty"`
	// Merchant’s order Id of the order for which refund is initiated
	OrderId *string `json:"order_id,omitempty"`
	// Merchant’s refund ID of the refund
	RefundId *string `json:"refund_id,omitempty"`
	// Type of object
	Entity *string `json:"entity,omitempty"`
	// Amount that is refunded
	RefundAmount *float32 `json:"refund_amount,omitempty"`
	// Currency of the refund amount
	RefundCurrency *string `json:"refund_currency,omitempty"`
	// Note added by merchant for the refund
	RefundNote *string `json:"refund_note,omitempty"`
	// This can be one of [\"SUCCESS\", \"PENDING\", \"CANCELLED\", \"ONHOLD\", \"FAILED\"]
	RefundStatus *string `json:"refund_status,omitempty"`
	// The bank reference number for refund
	RefundArn *string `json:"refund_arn,omitempty"`
	// Charges in INR for processing refund
	RefundCharge *float32 `json:"refund_charge,omitempty"`
	// Description of refund status
	StatusDescription *string `json:"status_description,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	RefundSplits []VendorSplit `json:"refund_splits,omitempty"`
	// This can be one of [\"PAYMENT_AUTO_REFUND\", \"MERCHANT_INITIATED\", \"UNRECONCILED_AUTO_REFUND\"]
	RefundType *string `json:"refund_type,omitempty"`
	// Method or speed of processing refund
	RefundMode *string `json:"refund_mode,omitempty"`
	// Time of refund creation
	CreatedAt *string `json:"created_at,omitempty"`
	// Time when refund was processed successfully
	ProcessedAt *string `json:"processed_at,omitempty"`
	RefundSpeed *RefundSpeed `json:"refund_speed,omitempty"`
}

// NewRefundEntity instantiates a new RefundEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundEntity() *RefundEntity {
	this := RefundEntity{}
	return &this
}

// NewRefundEntityWithDefaults instantiates a new RefundEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundEntityWithDefaults() *RefundEntity {
	this := RefundEntity{}
	return &this
}

// GetCfPaymentId returns the CfPaymentId field value if set, zero value otherwise.
func (o *RefundEntity) GetCfPaymentId() int32 {
	if o == nil || IsNil(o.CfPaymentId) {
		var ret int32
		return ret
	}
	return *o.CfPaymentId
}

// GetCfPaymentIdOk returns a tuple with the CfPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetCfPaymentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CfPaymentId) {
		return nil, false
	}
	return o.CfPaymentId, true
}

// HasCfPaymentId returns a boolean if a field has been set.
func (o *RefundEntity) HasCfPaymentId() bool {
	if o != nil && !IsNil(o.CfPaymentId) {
		return true
	}

	return false
}

// SetCfPaymentId gets a reference to the given int32 and assigns it to the CfPaymentId field.
func (o *RefundEntity) SetCfPaymentId(v int32) {
	o.CfPaymentId = &v
}

// GetCfRefundId returns the CfRefundId field value if set, zero value otherwise.
func (o *RefundEntity) GetCfRefundId() string {
	if o == nil || IsNil(o.CfRefundId) {
		var ret string
		return ret
	}
	return *o.CfRefundId
}

// GetCfRefundIdOk returns a tuple with the CfRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetCfRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.CfRefundId) {
		return nil, false
	}
	return o.CfRefundId, true
}

// HasCfRefundId returns a boolean if a field has been set.
func (o *RefundEntity) HasCfRefundId() bool {
	if o != nil && !IsNil(o.CfRefundId) {
		return true
	}

	return false
}

// SetCfRefundId gets a reference to the given string and assigns it to the CfRefundId field.
func (o *RefundEntity) SetCfRefundId(v string) {
	o.CfRefundId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *RefundEntity) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *RefundEntity) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *RefundEntity) SetOrderId(v string) {
	o.OrderId = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *RefundEntity) SetRefundId(v string) {
	o.RefundId = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *RefundEntity) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *RefundEntity) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *RefundEntity) SetEntity(v string) {
	o.Entity = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundAmount() float32 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret float32
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given float32 and assigns it to the RefundAmount field.
func (o *RefundEntity) SetRefundAmount(v float32) {
	o.RefundAmount = &v
}

// GetRefundCurrency returns the RefundCurrency field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundCurrency() string {
	if o == nil || IsNil(o.RefundCurrency) {
		var ret string
		return ret
	}
	return *o.RefundCurrency
}

// GetRefundCurrencyOk returns a tuple with the RefundCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.RefundCurrency) {
		return nil, false
	}
	return o.RefundCurrency, true
}

// HasRefundCurrency returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundCurrency() bool {
	if o != nil && !IsNil(o.RefundCurrency) {
		return true
	}

	return false
}

// SetRefundCurrency gets a reference to the given string and assigns it to the RefundCurrency field.
func (o *RefundEntity) SetRefundCurrency(v string) {
	o.RefundCurrency = &v
}

// GetRefundNote returns the RefundNote field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundNote() string {
	if o == nil || IsNil(o.RefundNote) {
		var ret string
		return ret
	}
	return *o.RefundNote
}

// GetRefundNoteOk returns a tuple with the RefundNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundNoteOk() (*string, bool) {
	if o == nil || IsNil(o.RefundNote) {
		return nil, false
	}
	return o.RefundNote, true
}

// HasRefundNote returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundNote() bool {
	if o != nil && !IsNil(o.RefundNote) {
		return true
	}

	return false
}

// SetRefundNote gets a reference to the given string and assigns it to the RefundNote field.
func (o *RefundEntity) SetRefundNote(v string) {
	o.RefundNote = &v
}

// GetRefundStatus returns the RefundStatus field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundStatus() string {
	if o == nil || IsNil(o.RefundStatus) {
		var ret string
		return ret
	}
	return *o.RefundStatus
}

// GetRefundStatusOk returns a tuple with the RefundStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RefundStatus) {
		return nil, false
	}
	return o.RefundStatus, true
}

// HasRefundStatus returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundStatus() bool {
	if o != nil && !IsNil(o.RefundStatus) {
		return true
	}

	return false
}

// SetRefundStatus gets a reference to the given string and assigns it to the RefundStatus field.
func (o *RefundEntity) SetRefundStatus(v string) {
	o.RefundStatus = &v
}

// GetRefundArn returns the RefundArn field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundArn() string {
	if o == nil || IsNil(o.RefundArn) {
		var ret string
		return ret
	}
	return *o.RefundArn
}

// GetRefundArnOk returns a tuple with the RefundArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundArnOk() (*string, bool) {
	if o == nil || IsNil(o.RefundArn) {
		return nil, false
	}
	return o.RefundArn, true
}

// HasRefundArn returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundArn() bool {
	if o != nil && !IsNil(o.RefundArn) {
		return true
	}

	return false
}

// SetRefundArn gets a reference to the given string and assigns it to the RefundArn field.
func (o *RefundEntity) SetRefundArn(v string) {
	o.RefundArn = &v
}

// GetRefundCharge returns the RefundCharge field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundCharge() float32 {
	if o == nil || IsNil(o.RefundCharge) {
		var ret float32
		return ret
	}
	return *o.RefundCharge
}

// GetRefundChargeOk returns a tuple with the RefundCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundCharge) {
		return nil, false
	}
	return o.RefundCharge, true
}

// HasRefundCharge returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundCharge() bool {
	if o != nil && !IsNil(o.RefundCharge) {
		return true
	}

	return false
}

// SetRefundCharge gets a reference to the given float32 and assigns it to the RefundCharge field.
func (o *RefundEntity) SetRefundCharge(v float32) {
	o.RefundCharge = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *RefundEntity) GetStatusDescription() string {
	if o == nil || IsNil(o.StatusDescription) {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDescription) {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *RefundEntity) HasStatusDescription() bool {
	if o != nil && !IsNil(o.StatusDescription) {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *RefundEntity) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundEntity) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundEntity) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RefundEntity) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RefundEntity) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRefundSplits returns the RefundSplits field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundSplits() []VendorSplit {
	if o == nil || IsNil(o.RefundSplits) {
		var ret []VendorSplit
		return ret
	}
	return o.RefundSplits
}

// GetRefundSplitsOk returns a tuple with the RefundSplits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundSplitsOk() ([]VendorSplit, bool) {
	if o == nil || IsNil(o.RefundSplits) {
		return nil, false
	}
	return o.RefundSplits, true
}

// HasRefundSplits returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundSplits() bool {
	if o != nil && !IsNil(o.RefundSplits) {
		return true
	}

	return false
}

// SetRefundSplits gets a reference to the given []VendorSplit and assigns it to the RefundSplits field.
func (o *RefundEntity) SetRefundSplits(v []VendorSplit) {
	o.RefundSplits = v
}

// GetRefundType returns the RefundType field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundType() string {
	if o == nil || IsNil(o.RefundType) {
		var ret string
		return ret
	}
	return *o.RefundType
}

// GetRefundTypeOk returns a tuple with the RefundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefundType) {
		return nil, false
	}
	return o.RefundType, true
}

// HasRefundType returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundType() bool {
	if o != nil && !IsNil(o.RefundType) {
		return true
	}

	return false
}

// SetRefundType gets a reference to the given string and assigns it to the RefundType field.
func (o *RefundEntity) SetRefundType(v string) {
	o.RefundType = &v
}

// GetRefundMode returns the RefundMode field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundMode() string {
	if o == nil || IsNil(o.RefundMode) {
		var ret string
		return ret
	}
	return *o.RefundMode
}

// GetRefundModeOk returns a tuple with the RefundMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundModeOk() (*string, bool) {
	if o == nil || IsNil(o.RefundMode) {
		return nil, false
	}
	return o.RefundMode, true
}

// HasRefundMode returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundMode() bool {
	if o != nil && !IsNil(o.RefundMode) {
		return true
	}

	return false
}

// SetRefundMode gets a reference to the given string and assigns it to the RefundMode field.
func (o *RefundEntity) SetRefundMode(v string) {
	o.RefundMode = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RefundEntity) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RefundEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RefundEntity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetProcessedAt returns the ProcessedAt field value if set, zero value otherwise.
func (o *RefundEntity) GetProcessedAt() string {
	if o == nil || IsNil(o.ProcessedAt) {
		var ret string
		return ret
	}
	return *o.ProcessedAt
}

// GetProcessedAtOk returns a tuple with the ProcessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetProcessedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessedAt) {
		return nil, false
	}
	return o.ProcessedAt, true
}

// HasProcessedAt returns a boolean if a field has been set.
func (o *RefundEntity) HasProcessedAt() bool {
	if o != nil && !IsNil(o.ProcessedAt) {
		return true
	}

	return false
}

// SetProcessedAt gets a reference to the given string and assigns it to the ProcessedAt field.
func (o *RefundEntity) SetProcessedAt(v string) {
	o.ProcessedAt = &v
}

// GetRefundSpeed returns the RefundSpeed field value if set, zero value otherwise.
func (o *RefundEntity) GetRefundSpeed() RefundSpeed {
	if o == nil || IsNil(o.RefundSpeed) {
		var ret RefundSpeed
		return ret
	}
	return *o.RefundSpeed
}

// GetRefundSpeedOk returns a tuple with the RefundSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundEntity) GetRefundSpeedOk() (*RefundSpeed, bool) {
	if o == nil || IsNil(o.RefundSpeed) {
		return nil, false
	}
	return o.RefundSpeed, true
}

// HasRefundSpeed returns a boolean if a field has been set.
func (o *RefundEntity) HasRefundSpeed() bool {
	if o != nil && !IsNil(o.RefundSpeed) {
		return true
	}

	return false
}

// SetRefundSpeed gets a reference to the given RefundSpeed and assigns it to the RefundSpeed field.
func (o *RefundEntity) SetRefundSpeed(v RefundSpeed) {
	o.RefundSpeed = &v
}

func (o RefundEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.CfRefundId) {
		toSerialize["cf_refund_id"] = o.CfRefundId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.RefundId) {
		toSerialize["refund_id"] = o.RefundId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.RefundCurrency) {
		toSerialize["refund_currency"] = o.RefundCurrency
	}
	if !IsNil(o.RefundNote) {
		toSerialize["refund_note"] = o.RefundNote
	}
	if !IsNil(o.RefundStatus) {
		toSerialize["refund_status"] = o.RefundStatus
	}
	if !IsNil(o.RefundArn) {
		toSerialize["refund_arn"] = o.RefundArn
	}
	if !IsNil(o.RefundCharge) {
		toSerialize["refund_charge"] = o.RefundCharge
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["status_description"] = o.StatusDescription
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RefundSplits) {
		toSerialize["refund_splits"] = o.RefundSplits
	}
	if !IsNil(o.RefundType) {
		toSerialize["refund_type"] = o.RefundType
	}
	if !IsNil(o.RefundMode) {
		toSerialize["refund_mode"] = o.RefundMode
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ProcessedAt) {
		toSerialize["processed_at"] = o.ProcessedAt
	}
	if !IsNil(o.RefundSpeed) {
		toSerialize["refund_speed"] = o.RefundSpeed
	}
	return toSerialize, nil
}

type NullableRefundEntity struct {
	value *RefundEntity
	isSet bool
}

func (v NullableRefundEntity) Get() *RefundEntity {
	return v.value
}

func (v *NullableRefundEntity) Set(val *RefundEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundEntity(val *RefundEntity) *NullableRefundEntity {
	return &NullableRefundEntity{value: val, isSet: true}
}

func (v NullableRefundEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


