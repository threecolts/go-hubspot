/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"encoding/json"
)

// checks if the AssociationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociationSpec{}

// AssociationSpec struct for AssociationSpec
type AssociationSpec struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeId   int32  `json:"associationTypeId"`
}

// NewAssociationSpec instantiates a new AssociationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationSpec(associationCategory string, associationTypeId int32) *AssociationSpec {
	this := AssociationSpec{}
	this.AssociationCategory = associationCategory
	this.AssociationTypeId = associationTypeId
	return &this
}

// NewAssociationSpecWithDefaults instantiates a new AssociationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationSpecWithDefaults() *AssociationSpec {
	this := AssociationSpec{}
	return &this
}

// GetAssociationCategory returns the AssociationCategory field value
func (o *AssociationSpec) GetAssociationCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociationCategory
}

// GetAssociationCategoryOk returns a tuple with the AssociationCategory field value
// and a boolean to check if the value has been set.
func (o *AssociationSpec) GetAssociationCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationCategory, true
}

// SetAssociationCategory sets field value
func (o *AssociationSpec) SetAssociationCategory(v string) {
	o.AssociationCategory = v
}

// GetAssociationTypeId returns the AssociationTypeId field value
func (o *AssociationSpec) GetAssociationTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssociationTypeId
}

// GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationSpec) GetAssociationTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationTypeId, true
}

// SetAssociationTypeId sets field value
func (o *AssociationSpec) SetAssociationTypeId(v int32) {
	o.AssociationTypeId = v
}

func (o AssociationSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associationCategory"] = o.AssociationCategory
	toSerialize["associationTypeId"] = o.AssociationTypeId
	return toSerialize, nil
}

type NullableAssociationSpec struct {
	value *AssociationSpec
	isSet bool
}

func (v NullableAssociationSpec) Get() *AssociationSpec {
	return v.value
}

func (v *NullableAssociationSpec) Set(val *AssociationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationSpec(val *AssociationSpec) *NullableAssociationSpec {
	return &NullableAssociationSpec{value: val, isSet: true}
}

func (v NullableAssociationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
