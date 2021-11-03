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

// CollectionResponsePropertyGroup struct for CollectionResponsePropertyGroup
type CollectionResponsePropertyGroup struct {
	Results []PropertyGroup `json:"results"`
	Paging  *Paging         `json:"paging,omitempty"`
}

// NewCollectionResponsePropertyGroup instantiates a new CollectionResponsePropertyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePropertyGroup(results []PropertyGroup) *CollectionResponsePropertyGroup {
	this := CollectionResponsePropertyGroup{}
	this.Results = results
	return &this
}

// NewCollectionResponsePropertyGroupWithDefaults instantiates a new CollectionResponsePropertyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePropertyGroupWithDefaults() *CollectionResponsePropertyGroup {
	this := CollectionResponsePropertyGroup{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponsePropertyGroup) GetResults() []PropertyGroup {
	if o == nil {
		var ret []PropertyGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePropertyGroup) GetResultsOk() (*[]PropertyGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePropertyGroup) SetResults(v []PropertyGroup) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponsePropertyGroup) GetPaging() Paging {
	if o == nil || o.Paging == nil {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponsePropertyGroup) GetPagingOk() (*Paging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponsePropertyGroup) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponsePropertyGroup) SetPaging(v Paging) {
	o.Paging = &v
}

func (o CollectionResponsePropertyGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponsePropertyGroup struct {
	value *CollectionResponsePropertyGroup
	isSet bool
}

func (v NullableCollectionResponsePropertyGroup) Get() *CollectionResponsePropertyGroup {
	return v.value
}

func (v *NullableCollectionResponsePropertyGroup) Set(val *CollectionResponsePropertyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePropertyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePropertyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePropertyGroup(val *CollectionResponsePropertyGroup) *NullableCollectionResponsePropertyGroup {
	return &NullableCollectionResponsePropertyGroup{value: val, isSet: true}
}

func (v NullableCollectionResponsePropertyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePropertyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
