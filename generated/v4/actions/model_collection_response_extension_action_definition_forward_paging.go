/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// checks if the CollectionResponseExtensionActionDefinitionForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseExtensionActionDefinitionForwardPaging{}

// CollectionResponseExtensionActionDefinitionForwardPaging struct for CollectionResponseExtensionActionDefinitionForwardPaging
type CollectionResponseExtensionActionDefinitionForwardPaging struct {
	Results []ExtensionActionDefinition `json:"results"`
	Paging  *ForwardPaging              `json:"paging,omitempty"`
}

// NewCollectionResponseExtensionActionDefinitionForwardPaging instantiates a new CollectionResponseExtensionActionDefinitionForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseExtensionActionDefinitionForwardPaging(results []ExtensionActionDefinition) *CollectionResponseExtensionActionDefinitionForwardPaging {
	this := CollectionResponseExtensionActionDefinitionForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseExtensionActionDefinitionForwardPagingWithDefaults instantiates a new CollectionResponseExtensionActionDefinitionForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseExtensionActionDefinitionForwardPagingWithDefaults() *CollectionResponseExtensionActionDefinitionForwardPaging {
	this := CollectionResponseExtensionActionDefinitionForwardPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) GetResults() []ExtensionActionDefinition {
	if o == nil {
		var ret []ExtensionActionDefinition
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) GetResultsOk() ([]ExtensionActionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) SetResults(v []ExtensionActionDefinition) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseExtensionActionDefinitionForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseExtensionActionDefinitionForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseExtensionActionDefinitionForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableCollectionResponseExtensionActionDefinitionForwardPaging struct {
	value *CollectionResponseExtensionActionDefinitionForwardPaging
	isSet bool
}

func (v NullableCollectionResponseExtensionActionDefinitionForwardPaging) Get() *CollectionResponseExtensionActionDefinitionForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseExtensionActionDefinitionForwardPaging) Set(val *CollectionResponseExtensionActionDefinitionForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseExtensionActionDefinitionForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseExtensionActionDefinitionForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseExtensionActionDefinitionForwardPaging(val *CollectionResponseExtensionActionDefinitionForwardPaging) *NullableCollectionResponseExtensionActionDefinitionForwardPaging {
	return &NullableCollectionResponseExtensionActionDefinitionForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseExtensionActionDefinitionForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseExtensionActionDefinitionForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
