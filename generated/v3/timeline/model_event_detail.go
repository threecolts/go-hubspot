/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"encoding/json"
)

// checks if the EventDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDetail{}

// EventDetail The details Markdown rendered as HTML.
type EventDetail struct {
	// The details Markdown rendered as HTML.
	Details string `json:"details"`
}

// NewEventDetail instantiates a new EventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDetail(details string) *EventDetail {
	this := EventDetail{}
	this.Details = details
	return &this
}

// NewEventDetailWithDefaults instantiates a new EventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDetailWithDefaults() *EventDetail {
	this := EventDetail{}
	return &this
}

// GetDetails returns the Details field value
func (o *EventDetail) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *EventDetail) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *EventDetail) SetDetails(v string) {
	o.Details = v
}

func (o EventDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableEventDetail struct {
	value *EventDetail
	isSet bool
}

func (v NullableEventDetail) Get() *EventDetail {
	return v.value
}

func (v *NullableEventDetail) Set(val *EventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDetail(val *EventDetail) *NullableEventDetail {
	return &NullableEventDetail{value: val, isSet: true}
}

func (v NullableEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
