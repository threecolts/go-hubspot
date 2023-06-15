/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
)

// checks if the UpdateLanguagesRequestVNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLanguagesRequestVNext{}

// UpdateLanguagesRequestVNext Request object for updating languages within a multi-language group.
type UpdateLanguagesRequestVNext struct {
	// ID of the primary object in the multi-language group.
	PrimaryId string `json:"primaryId"`
	// Map of object IDs to associated languages of object in the multi-language group.
	Languages map[string]string `json:"languages"`
}

// NewUpdateLanguagesRequestVNext instantiates a new UpdateLanguagesRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLanguagesRequestVNext(primaryId string, languages map[string]string) *UpdateLanguagesRequestVNext {
	this := UpdateLanguagesRequestVNext{}
	this.PrimaryId = primaryId
	this.Languages = languages
	return &this
}

// NewUpdateLanguagesRequestVNextWithDefaults instantiates a new UpdateLanguagesRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLanguagesRequestVNextWithDefaults() *UpdateLanguagesRequestVNext {
	this := UpdateLanguagesRequestVNext{}
	return &this
}

// GetPrimaryId returns the PrimaryId field value
func (o *UpdateLanguagesRequestVNext) GetPrimaryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryId
}

// GetPrimaryIdOk returns a tuple with the PrimaryId field value
// and a boolean to check if the value has been set.
func (o *UpdateLanguagesRequestVNext) GetPrimaryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryId, true
}

// SetPrimaryId sets field value
func (o *UpdateLanguagesRequestVNext) SetPrimaryId(v string) {
	o.PrimaryId = v
}

// GetLanguages returns the Languages field value
func (o *UpdateLanguagesRequestVNext) GetLanguages() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *UpdateLanguagesRequestVNext) GetLanguagesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Languages, true
}

// SetLanguages sets field value
func (o *UpdateLanguagesRequestVNext) SetLanguages(v map[string]string) {
	o.Languages = v
}

func (o UpdateLanguagesRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLanguagesRequestVNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primaryId"] = o.PrimaryId
	toSerialize["languages"] = o.Languages
	return toSerialize, nil
}

type NullableUpdateLanguagesRequestVNext struct {
	value *UpdateLanguagesRequestVNext
	isSet bool
}

func (v NullableUpdateLanguagesRequestVNext) Get() *UpdateLanguagesRequestVNext {
	return v.value
}

func (v *NullableUpdateLanguagesRequestVNext) Set(val *UpdateLanguagesRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLanguagesRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLanguagesRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLanguagesRequestVNext(val *UpdateLanguagesRequestVNext) *NullableUpdateLanguagesRequestVNext {
	return &NullableUpdateLanguagesRequestVNext{value: val, isSet: true}
}

func (v NullableUpdateLanguagesRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLanguagesRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
