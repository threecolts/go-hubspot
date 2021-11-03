/*
CRM cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_extensions

import (
	"encoding/json"
)

// CardObjectTypeBody struct for CardObjectTypeBody
type CardObjectTypeBody struct {
	// A CRM object type where this card should be displayed.
	Name string `json:"name"`
	// An array of properties that should be sent to this card's target URL when the data fetch request is made. Must be valid properties for the corresponding CRM object type.
	PropertiesToSend []string `json:"propertiesToSend"`
}

// NewCardObjectTypeBody instantiates a new CardObjectTypeBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardObjectTypeBody(name string, propertiesToSend []string) *CardObjectTypeBody {
	this := CardObjectTypeBody{}
	this.Name = name
	this.PropertiesToSend = propertiesToSend
	return &this
}

// NewCardObjectTypeBodyWithDefaults instantiates a new CardObjectTypeBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardObjectTypeBodyWithDefaults() *CardObjectTypeBody {
	this := CardObjectTypeBody{}
	return &this
}

// GetName returns the Name field value
func (o *CardObjectTypeBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CardObjectTypeBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CardObjectTypeBody) SetName(v string) {
	o.Name = v
}

// GetPropertiesToSend returns the PropertiesToSend field value
func (o *CardObjectTypeBody) GetPropertiesToSend() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertiesToSend
}

// GetPropertiesToSendOk returns a tuple with the PropertiesToSend field value
// and a boolean to check if the value has been set.
func (o *CardObjectTypeBody) GetPropertiesToSendOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertiesToSend, true
}

// SetPropertiesToSend sets field value
func (o *CardObjectTypeBody) SetPropertiesToSend(v []string) {
	o.PropertiesToSend = v
}

func (o CardObjectTypeBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["propertiesToSend"] = o.PropertiesToSend
	}
	return json.Marshal(toSerialize)
}

type NullableCardObjectTypeBody struct {
	value *CardObjectTypeBody
	isSet bool
}

func (v NullableCardObjectTypeBody) Get() *CardObjectTypeBody {
	return v.value
}

func (v *NullableCardObjectTypeBody) Set(val *CardObjectTypeBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCardObjectTypeBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCardObjectTypeBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardObjectTypeBody(val *CardObjectTypeBody) *NullableCardObjectTypeBody {
	return &NullableCardObjectTypeBody{value: val, isSet: true}
}

func (v NullableCardObjectTypeBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardObjectTypeBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
