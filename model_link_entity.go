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

// checks if the LinkEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkEntity{}

// LinkEntity Payment link success creation response object
type LinkEntity struct {
	CfLinkId *int32 `json:"cf_link_id,omitempty"`
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
	LinkMeta *LinkMetaEntity `json:"link_meta,omitempty"`
	LinkUrl *string `json:"link_url,omitempty"`
	LinkExpiryTime *string `json:"link_expiry_time,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	LinkNotes *map[string]string `json:"link_notes,omitempty"`
	LinkAutoReminders *bool `json:"link_auto_reminders,omitempty"`
	LinkNotify *LinkNotifyEntity `json:"link_notify,omitempty"`
}

// NewLinkEntity instantiates a new LinkEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkEntity() *LinkEntity {
	this := LinkEntity{}
	return &this
}

// NewLinkEntityWithDefaults instantiates a new LinkEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkEntityWithDefaults() *LinkEntity {
	this := LinkEntity{}
	return &this
}

// GetCfLinkId returns the CfLinkId field value if set, zero value otherwise.
func (o *LinkEntity) GetCfLinkId() int32 {
	if o == nil || IsNil(o.CfLinkId) {
		var ret int32
		return ret
	}
	return *o.CfLinkId
}

// GetCfLinkIdOk returns a tuple with the CfLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetCfLinkIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CfLinkId) {
		return nil, false
	}
	return o.CfLinkId, true
}

// HasCfLinkId returns a boolean if a field has been set.
func (o *LinkEntity) HasCfLinkId() bool {
	if o != nil && !IsNil(o.CfLinkId) {
		return true
	}

	return false
}

// SetCfLinkId gets a reference to the given int32 and assigns it to the CfLinkId field.
func (o *LinkEntity) SetCfLinkId(v int32) {
	o.CfLinkId = &v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *LinkEntity) SetLinkId(v string) {
	o.LinkId = &v
}

// GetLinkStatus returns the LinkStatus field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkStatus() string {
	if o == nil || IsNil(o.LinkStatus) {
		var ret string
		return ret
	}
	return *o.LinkStatus
}

// GetLinkStatusOk returns a tuple with the LinkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LinkStatus) {
		return nil, false
	}
	return o.LinkStatus, true
}

// HasLinkStatus returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkStatus() bool {
	if o != nil && !IsNil(o.LinkStatus) {
		return true
	}

	return false
}

// SetLinkStatus gets a reference to the given string and assigns it to the LinkStatus field.
func (o *LinkEntity) SetLinkStatus(v string) {
	o.LinkStatus = &v
}

// GetLinkCurrency returns the LinkCurrency field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkCurrency() string {
	if o == nil || IsNil(o.LinkCurrency) {
		var ret string
		return ret
	}
	return *o.LinkCurrency
}

// GetLinkCurrencyOk returns a tuple with the LinkCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.LinkCurrency) {
		return nil, false
	}
	return o.LinkCurrency, true
}

// HasLinkCurrency returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkCurrency() bool {
	if o != nil && !IsNil(o.LinkCurrency) {
		return true
	}

	return false
}

// SetLinkCurrency gets a reference to the given string and assigns it to the LinkCurrency field.
func (o *LinkEntity) SetLinkCurrency(v string) {
	o.LinkCurrency = &v
}

// GetLinkAmount returns the LinkAmount field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkAmount() float32 {
	if o == nil || IsNil(o.LinkAmount) {
		var ret float32
		return ret
	}
	return *o.LinkAmount
}

// GetLinkAmountOk returns a tuple with the LinkAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LinkAmount) {
		return nil, false
	}
	return o.LinkAmount, true
}

// HasLinkAmount returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkAmount() bool {
	if o != nil && !IsNil(o.LinkAmount) {
		return true
	}

	return false
}

// SetLinkAmount gets a reference to the given float32 and assigns it to the LinkAmount field.
func (o *LinkEntity) SetLinkAmount(v float32) {
	o.LinkAmount = &v
}

// GetLinkAmountPaid returns the LinkAmountPaid field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkAmountPaid() float32 {
	if o == nil || IsNil(o.LinkAmountPaid) {
		var ret float32
		return ret
	}
	return *o.LinkAmountPaid
}

// GetLinkAmountPaidOk returns a tuple with the LinkAmountPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkAmountPaidOk() (*float32, bool) {
	if o == nil || IsNil(o.LinkAmountPaid) {
		return nil, false
	}
	return o.LinkAmountPaid, true
}

// HasLinkAmountPaid returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkAmountPaid() bool {
	if o != nil && !IsNil(o.LinkAmountPaid) {
		return true
	}

	return false
}

// SetLinkAmountPaid gets a reference to the given float32 and assigns it to the LinkAmountPaid field.
func (o *LinkEntity) SetLinkAmountPaid(v float32) {
	o.LinkAmountPaid = &v
}

// GetLinkPartialPayments returns the LinkPartialPayments field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkPartialPayments() bool {
	if o == nil || IsNil(o.LinkPartialPayments) {
		var ret bool
		return ret
	}
	return *o.LinkPartialPayments
}

// GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkPartialPaymentsOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkPartialPayments) {
		return nil, false
	}
	return o.LinkPartialPayments, true
}

// HasLinkPartialPayments returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkPartialPayments() bool {
	if o != nil && !IsNil(o.LinkPartialPayments) {
		return true
	}

	return false
}

// SetLinkPartialPayments gets a reference to the given bool and assigns it to the LinkPartialPayments field.
func (o *LinkEntity) SetLinkPartialPayments(v bool) {
	o.LinkPartialPayments = &v
}

// GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkMinimumPartialAmount() float32 {
	if o == nil || IsNil(o.LinkMinimumPartialAmount) {
		var ret float32
		return ret
	}
	return *o.LinkMinimumPartialAmount
}

// GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkMinimumPartialAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LinkMinimumPartialAmount) {
		return nil, false
	}
	return o.LinkMinimumPartialAmount, true
}

// HasLinkMinimumPartialAmount returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkMinimumPartialAmount() bool {
	if o != nil && !IsNil(o.LinkMinimumPartialAmount) {
		return true
	}

	return false
}

// SetLinkMinimumPartialAmount gets a reference to the given float32 and assigns it to the LinkMinimumPartialAmount field.
func (o *LinkEntity) SetLinkMinimumPartialAmount(v float32) {
	o.LinkMinimumPartialAmount = &v
}

// GetLinkPurpose returns the LinkPurpose field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkPurpose() string {
	if o == nil || IsNil(o.LinkPurpose) {
		var ret string
		return ret
	}
	return *o.LinkPurpose
}

// GetLinkPurposeOk returns a tuple with the LinkPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkPurpose) {
		return nil, false
	}
	return o.LinkPurpose, true
}

// HasLinkPurpose returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkPurpose() bool {
	if o != nil && !IsNil(o.LinkPurpose) {
		return true
	}

	return false
}

// SetLinkPurpose gets a reference to the given string and assigns it to the LinkPurpose field.
func (o *LinkEntity) SetLinkPurpose(v string) {
	o.LinkPurpose = &v
}

// GetLinkCreatedAt returns the LinkCreatedAt field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkCreatedAt() string {
	if o == nil || IsNil(o.LinkCreatedAt) {
		var ret string
		return ret
	}
	return *o.LinkCreatedAt
}

// GetLinkCreatedAtOk returns a tuple with the LinkCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LinkCreatedAt) {
		return nil, false
	}
	return o.LinkCreatedAt, true
}

// HasLinkCreatedAt returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkCreatedAt() bool {
	if o != nil && !IsNil(o.LinkCreatedAt) {
		return true
	}

	return false
}

// SetLinkCreatedAt gets a reference to the given string and assigns it to the LinkCreatedAt field.
func (o *LinkEntity) SetLinkCreatedAt(v string) {
	o.LinkCreatedAt = &v
}

// GetCustomerDetails returns the CustomerDetails field value if set, zero value otherwise.
func (o *LinkEntity) GetCustomerDetails() LinkCustomerDetailsEntity {
	if o == nil || IsNil(o.CustomerDetails) {
		var ret LinkCustomerDetailsEntity
		return ret
	}
	return *o.CustomerDetails
}

// GetCustomerDetailsOk returns a tuple with the CustomerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetCustomerDetailsOk() (*LinkCustomerDetailsEntity, bool) {
	if o == nil || IsNil(o.CustomerDetails) {
		return nil, false
	}
	return o.CustomerDetails, true
}

// HasCustomerDetails returns a boolean if a field has been set.
func (o *LinkEntity) HasCustomerDetails() bool {
	if o != nil && !IsNil(o.CustomerDetails) {
		return true
	}

	return false
}

// SetCustomerDetails gets a reference to the given LinkCustomerDetailsEntity and assigns it to the CustomerDetails field.
func (o *LinkEntity) SetCustomerDetails(v LinkCustomerDetailsEntity) {
	o.CustomerDetails = &v
}

// GetLinkMeta returns the LinkMeta field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkMeta() LinkMetaEntity {
	if o == nil || IsNil(o.LinkMeta) {
		var ret LinkMetaEntity
		return ret
	}
	return *o.LinkMeta
}

// GetLinkMetaOk returns a tuple with the LinkMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkMetaOk() (*LinkMetaEntity, bool) {
	if o == nil || IsNil(o.LinkMeta) {
		return nil, false
	}
	return o.LinkMeta, true
}

// HasLinkMeta returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkMeta() bool {
	if o != nil && !IsNil(o.LinkMeta) {
		return true
	}

	return false
}

// SetLinkMeta gets a reference to the given LinkMetaEntity and assigns it to the LinkMeta field.
func (o *LinkEntity) SetLinkMeta(v LinkMetaEntity) {
	o.LinkMeta = &v
}

// GetLinkUrl returns the LinkUrl field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkUrl() string {
	if o == nil || IsNil(o.LinkUrl) {
		var ret string
		return ret
	}
	return *o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkUrl) {
		return nil, false
	}
	return o.LinkUrl, true
}

// HasLinkUrl returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkUrl() bool {
	if o != nil && !IsNil(o.LinkUrl) {
		return true
	}

	return false
}

// SetLinkUrl gets a reference to the given string and assigns it to the LinkUrl field.
func (o *LinkEntity) SetLinkUrl(v string) {
	o.LinkUrl = &v
}

// GetLinkExpiryTime returns the LinkExpiryTime field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkExpiryTime() string {
	if o == nil || IsNil(o.LinkExpiryTime) {
		var ret string
		return ret
	}
	return *o.LinkExpiryTime
}

// GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkExpiryTime) {
		return nil, false
	}
	return o.LinkExpiryTime, true
}

// HasLinkExpiryTime returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkExpiryTime() bool {
	if o != nil && !IsNil(o.LinkExpiryTime) {
		return true
	}

	return false
}

// SetLinkExpiryTime gets a reference to the given string and assigns it to the LinkExpiryTime field.
func (o *LinkEntity) SetLinkExpiryTime(v string) {
	o.LinkExpiryTime = &v
}

// GetLinkNotes returns the LinkNotes field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkNotes() map[string]string {
	if o == nil || IsNil(o.LinkNotes) {
		var ret map[string]string
		return ret
	}
	return *o.LinkNotes
}

// GetLinkNotesOk returns a tuple with the LinkNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkNotesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.LinkNotes) {
		return nil, false
	}
	return o.LinkNotes, true
}

// HasLinkNotes returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkNotes() bool {
	if o != nil && !IsNil(o.LinkNotes) {
		return true
	}

	return false
}

// SetLinkNotes gets a reference to the given map[string]string and assigns it to the LinkNotes field.
func (o *LinkEntity) SetLinkNotes(v map[string]string) {
	o.LinkNotes = &v
}

// GetLinkAutoReminders returns the LinkAutoReminders field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkAutoReminders() bool {
	if o == nil || IsNil(o.LinkAutoReminders) {
		var ret bool
		return ret
	}
	return *o.LinkAutoReminders
}

// GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkAutoRemindersOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkAutoReminders) {
		return nil, false
	}
	return o.LinkAutoReminders, true
}

// HasLinkAutoReminders returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkAutoReminders() bool {
	if o != nil && !IsNil(o.LinkAutoReminders) {
		return true
	}

	return false
}

// SetLinkAutoReminders gets a reference to the given bool and assigns it to the LinkAutoReminders field.
func (o *LinkEntity) SetLinkAutoReminders(v bool) {
	o.LinkAutoReminders = &v
}

// GetLinkNotify returns the LinkNotify field value if set, zero value otherwise.
func (o *LinkEntity) GetLinkNotify() LinkNotifyEntity {
	if o == nil || IsNil(o.LinkNotify) {
		var ret LinkNotifyEntity
		return ret
	}
	return *o.LinkNotify
}

// GetLinkNotifyOk returns a tuple with the LinkNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkEntity) GetLinkNotifyOk() (*LinkNotifyEntity, bool) {
	if o == nil || IsNil(o.LinkNotify) {
		return nil, false
	}
	return o.LinkNotify, true
}

// HasLinkNotify returns a boolean if a field has been set.
func (o *LinkEntity) HasLinkNotify() bool {
	if o != nil && !IsNil(o.LinkNotify) {
		return true
	}

	return false
}

// SetLinkNotify gets a reference to the given LinkNotifyEntity and assigns it to the LinkNotify field.
func (o *LinkEntity) SetLinkNotify(v LinkNotifyEntity) {
	o.LinkNotify = &v
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

type NullableLinkEntity struct {
	value *LinkEntity
	isSet bool
}

func (v NullableLinkEntity) Get() *LinkEntity {
	return v.value
}

func (v *NullableLinkEntity) Set(val *LinkEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkEntity(val *LinkEntity) *NullableLinkEntity {
	return &NullableLinkEntity{value: val, isSet: true}
}

func (v NullableLinkEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


