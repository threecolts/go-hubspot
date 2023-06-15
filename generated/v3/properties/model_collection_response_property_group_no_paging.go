/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
)

// checks if the CollectionResponsePropertyGroupNoPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponsePropertyGroupNoPaging{}

// CollectionResponsePropertyGroupNoPaging struct for CollectionResponsePropertyGroupNoPaging
type CollectionResponsePropertyGroupNoPaging struct {
	Results []PropertyGroup `json:"results"`
}

// NewCollectionResponsePropertyGroupNoPaging instantiates a new CollectionResponsePropertyGroupNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePropertyGroupNoPaging(results []PropertyGroup) *CollectionResponsePropertyGroupNoPaging {
	this := CollectionResponsePropertyGroupNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponsePropertyGroupNoPagingWithDefaults instantiates a new CollectionResponsePropertyGroupNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePropertyGroupNoPagingWithDefaults() *CollectionResponsePropertyGroupNoPaging {
	this := CollectionResponsePropertyGroupNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponsePropertyGroupNoPaging) GetResults() []PropertyGroup {
	if o == nil {
		var ret []PropertyGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePropertyGroupNoPaging) GetResultsOk() ([]PropertyGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePropertyGroupNoPaging) SetResults(v []PropertyGroup) {
	o.Results = v
}

func (o CollectionResponsePropertyGroupNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponsePropertyGroupNoPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableCollectionResponsePropertyGroupNoPaging struct {
	value *CollectionResponsePropertyGroupNoPaging
	isSet bool
}

func (v NullableCollectionResponsePropertyGroupNoPaging) Get() *CollectionResponsePropertyGroupNoPaging {
	return v.value
}

func (v *NullableCollectionResponsePropertyGroupNoPaging) Set(val *CollectionResponsePropertyGroupNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePropertyGroupNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePropertyGroupNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePropertyGroupNoPaging(val *CollectionResponsePropertyGroupNoPaging) *NullableCollectionResponsePropertyGroupNoPaging {
	return &NullableCollectionResponsePropertyGroupNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponsePropertyGroupNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePropertyGroupNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
