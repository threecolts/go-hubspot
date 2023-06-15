/*
Associations

Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"encoding/json"
)

// checks if the BatchInputPublicObjectId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputPublicObjectId{}

// BatchInputPublicObjectId struct for BatchInputPublicObjectId
type BatchInputPublicObjectId struct {
	Inputs []PublicObjectId `json:"inputs"`
}

// NewBatchInputPublicObjectId instantiates a new BatchInputPublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputPublicObjectId(inputs []PublicObjectId) *BatchInputPublicObjectId {
	this := BatchInputPublicObjectId{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputPublicObjectIdWithDefaults instantiates a new BatchInputPublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputPublicObjectIdWithDefaults() *BatchInputPublicObjectId {
	this := BatchInputPublicObjectId{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputPublicObjectId) GetInputs() []PublicObjectId {
	if o == nil {
		var ret []PublicObjectId
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputPublicObjectId) GetInputsOk() ([]PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputPublicObjectId) SetInputs(v []PublicObjectId) {
	o.Inputs = v
}

func (o BatchInputPublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputPublicObjectId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputPublicObjectId struct {
	value *BatchInputPublicObjectId
	isSet bool
}

func (v NullableBatchInputPublicObjectId) Get() *BatchInputPublicObjectId {
	return v.value
}

func (v *NullableBatchInputPublicObjectId) Set(val *BatchInputPublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputPublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputPublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputPublicObjectId(val *BatchInputPublicObjectId) *NullableBatchInputPublicObjectId {
	return &NullableBatchInputPublicObjectId{value: val, isSet: true}
}

func (v NullableBatchInputPublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputPublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
