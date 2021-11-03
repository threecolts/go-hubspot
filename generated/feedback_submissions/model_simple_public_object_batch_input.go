/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"encoding/json"
)

// SimplePublicObjectBatchInput struct for SimplePublicObjectBatchInput
type SimplePublicObjectBatchInput struct {
	Properties map[string]string `json:"properties"`
	Id         string            `json:"id"`
}

// NewSimplePublicObjectBatchInput instantiates a new SimplePublicObjectBatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectBatchInput(properties map[string]string, id string) *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	this.Properties = properties
	this.Id = id
	return &this
}

// NewSimplePublicObjectBatchInputWithDefaults instantiates a new SimplePublicObjectBatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectBatchInputWithDefaults() *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	return &this
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectBatchInput) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectBatchInput) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetId returns the Id field value
func (o *SimplePublicObjectBatchInput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectBatchInput) SetId(v string) {
	o.Id = v
}

func (o SimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSimplePublicObjectBatchInput struct {
	value *SimplePublicObjectBatchInput
	isSet bool
}

func (v NullableSimplePublicObjectBatchInput) Get() *SimplePublicObjectBatchInput {
	return v.value
}

func (v *NullableSimplePublicObjectBatchInput) Set(val *SimplePublicObjectBatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectBatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectBatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectBatchInput(val *SimplePublicObjectBatchInput) *NullableSimplePublicObjectBatchInput {
	return &NullableSimplePublicObjectBatchInput{value: val, isSet: true}
}

func (v NullableSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectBatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
