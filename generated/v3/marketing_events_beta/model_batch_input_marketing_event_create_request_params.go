/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// checks if the BatchInputMarketingEventCreateRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputMarketingEventCreateRequestParams{}

// BatchInputMarketingEventCreateRequestParams struct for BatchInputMarketingEventCreateRequestParams
type BatchInputMarketingEventCreateRequestParams struct {
	Inputs []MarketingEventCreateRequestParams `json:"inputs"`
}

// NewBatchInputMarketingEventCreateRequestParams instantiates a new BatchInputMarketingEventCreateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputMarketingEventCreateRequestParams(inputs []MarketingEventCreateRequestParams) *BatchInputMarketingEventCreateRequestParams {
	this := BatchInputMarketingEventCreateRequestParams{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputMarketingEventCreateRequestParamsWithDefaults instantiates a new BatchInputMarketingEventCreateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputMarketingEventCreateRequestParamsWithDefaults() *BatchInputMarketingEventCreateRequestParams {
	this := BatchInputMarketingEventCreateRequestParams{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputMarketingEventCreateRequestParams) GetInputs() []MarketingEventCreateRequestParams {
	if o == nil {
		var ret []MarketingEventCreateRequestParams
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputMarketingEventCreateRequestParams) GetInputsOk() ([]MarketingEventCreateRequestParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputMarketingEventCreateRequestParams) SetInputs(v []MarketingEventCreateRequestParams) {
	o.Inputs = v
}

func (o BatchInputMarketingEventCreateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputMarketingEventCreateRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputMarketingEventCreateRequestParams struct {
	value *BatchInputMarketingEventCreateRequestParams
	isSet bool
}

func (v NullableBatchInputMarketingEventCreateRequestParams) Get() *BatchInputMarketingEventCreateRequestParams {
	return v.value
}

func (v *NullableBatchInputMarketingEventCreateRequestParams) Set(val *BatchInputMarketingEventCreateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputMarketingEventCreateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputMarketingEventCreateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputMarketingEventCreateRequestParams(val *BatchInputMarketingEventCreateRequestParams) *NullableBatchInputMarketingEventCreateRequestParams {
	return &NullableBatchInputMarketingEventCreateRequestParams{value: val, isSet: true}
}

func (v NullableBatchInputMarketingEventCreateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputMarketingEventCreateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
