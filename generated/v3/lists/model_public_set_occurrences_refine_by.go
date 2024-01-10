/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicSetOccurrencesRefineBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSetOccurrencesRefineBy{}

// PublicSetOccurrencesRefineBy struct for PublicSetOccurrencesRefineBy
type PublicSetOccurrencesRefineBy struct {
	Type                 string `json:"type"`
	SetType_             string `json:"setType"`
	AdditionalProperties map[string]interface{}
}

type _PublicSetOccurrencesRefineBy PublicSetOccurrencesRefineBy

// NewPublicSetOccurrencesRefineBy instantiates a new PublicSetOccurrencesRefineBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSetOccurrencesRefineBy(type_ string, setType string) *PublicSetOccurrencesRefineBy {
	this := PublicSetOccurrencesRefineBy{}
	this.Type = type_
	this.SetType_ = setType
	return &this
}

// NewPublicSetOccurrencesRefineByWithDefaults instantiates a new PublicSetOccurrencesRefineBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSetOccurrencesRefineByWithDefaults() *PublicSetOccurrencesRefineBy {
	this := PublicSetOccurrencesRefineBy{}
	var type_ string = "SET_OCCURRENCES"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PublicSetOccurrencesRefineBy) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PublicSetOccurrencesRefineBy) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PublicSetOccurrencesRefineBy) SetType(v string) {
	o.Type = v
}

// GetSetType_ returns the SetType_ field value
func (o *PublicSetOccurrencesRefineBy) GetSetType_() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetType_
}

// GetSetType_Ok returns a tuple with the SetType_ field value
// and a boolean to check if the value has been set.
func (o *PublicSetOccurrencesRefineBy) GetSetType_Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetType_, true
}

// SetSetType_ sets field value
func (o *PublicSetOccurrencesRefineBy) SetSetType_(v string) {
	o.SetType_ = v
}

func (o PublicSetOccurrencesRefineBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSetOccurrencesRefineBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["setType"] = o.SetType_

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicSetOccurrencesRefineBy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"setType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicSetOccurrencesRefineBy := _PublicSetOccurrencesRefineBy{}

	err = json.Unmarshal(data, &varPublicSetOccurrencesRefineBy)

	if err != nil {
		return err
	}

	*o = PublicSetOccurrencesRefineBy(varPublicSetOccurrencesRefineBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "setType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicSetOccurrencesRefineBy struct {
	value *PublicSetOccurrencesRefineBy
	isSet bool
}

func (v NullablePublicSetOccurrencesRefineBy) Get() *PublicSetOccurrencesRefineBy {
	return v.value
}

func (v *NullablePublicSetOccurrencesRefineBy) Set(val *PublicSetOccurrencesRefineBy) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSetOccurrencesRefineBy) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSetOccurrencesRefineBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSetOccurrencesRefineBy(val *PublicSetOccurrencesRefineBy) *NullablePublicSetOccurrencesRefineBy {
	return &NullablePublicSetOccurrencesRefineBy{value: val, isSet: true}
}

func (v NullablePublicSetOccurrencesRefineBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSetOccurrencesRefineBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
