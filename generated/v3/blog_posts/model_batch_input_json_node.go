/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
)

// checks if the BatchInputJsonNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputJsonNode{}

// BatchInputJsonNode Wrapper for providing an array of JSON nodes as inputs.
type BatchInputJsonNode struct {
	// JSON nodes to input.
	Inputs []map[string]interface{} `json:"inputs"`
}

// NewBatchInputJsonNode instantiates a new BatchInputJsonNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputJsonNode(inputs []map[string]interface{}) *BatchInputJsonNode {
	this := BatchInputJsonNode{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputJsonNodeWithDefaults instantiates a new BatchInputJsonNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputJsonNodeWithDefaults() *BatchInputJsonNode {
	this := BatchInputJsonNode{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputJsonNode) GetInputs() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputJsonNode) GetInputsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputJsonNode) SetInputs(v []map[string]interface{}) {
	o.Inputs = v
}

func (o BatchInputJsonNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputJsonNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputJsonNode struct {
	value *BatchInputJsonNode
	isSet bool
}

func (v NullableBatchInputJsonNode) Get() *BatchInputJsonNode {
	return v.value
}

func (v *NullableBatchInputJsonNode) Set(val *BatchInputJsonNode) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputJsonNode) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputJsonNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputJsonNode(val *BatchInputJsonNode) *NullableBatchInputJsonNode {
	return &NullableBatchInputJsonNode{value: val, isSet: true}
}

func (v NullableBatchInputJsonNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputJsonNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
