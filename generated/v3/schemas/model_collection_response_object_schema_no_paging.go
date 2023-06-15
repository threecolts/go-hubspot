/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

import (
	"encoding/json"
)

// checks if the CollectionResponseObjectSchemaNoPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseObjectSchemaNoPaging{}

// CollectionResponseObjectSchemaNoPaging struct for CollectionResponseObjectSchemaNoPaging
type CollectionResponseObjectSchemaNoPaging struct {
	Results []ObjectSchema `json:"results"`
}

// NewCollectionResponseObjectSchemaNoPaging instantiates a new CollectionResponseObjectSchemaNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseObjectSchemaNoPaging(results []ObjectSchema) *CollectionResponseObjectSchemaNoPaging {
	this := CollectionResponseObjectSchemaNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseObjectSchemaNoPagingWithDefaults instantiates a new CollectionResponseObjectSchemaNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseObjectSchemaNoPagingWithDefaults() *CollectionResponseObjectSchemaNoPaging {
	this := CollectionResponseObjectSchemaNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseObjectSchemaNoPaging) GetResults() []ObjectSchema {
	if o == nil {
		var ret []ObjectSchema
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseObjectSchemaNoPaging) GetResultsOk() ([]ObjectSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseObjectSchemaNoPaging) SetResults(v []ObjectSchema) {
	o.Results = v
}

func (o CollectionResponseObjectSchemaNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseObjectSchemaNoPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableCollectionResponseObjectSchemaNoPaging struct {
	value *CollectionResponseObjectSchemaNoPaging
	isSet bool
}

func (v NullableCollectionResponseObjectSchemaNoPaging) Get() *CollectionResponseObjectSchemaNoPaging {
	return v.value
}

func (v *NullableCollectionResponseObjectSchemaNoPaging) Set(val *CollectionResponseObjectSchemaNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseObjectSchemaNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseObjectSchemaNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseObjectSchemaNoPaging(val *CollectionResponseObjectSchemaNoPaging) *NullableCollectionResponseObjectSchemaNoPaging {
	return &NullableCollectionResponseObjectSchemaNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseObjectSchemaNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseObjectSchemaNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
