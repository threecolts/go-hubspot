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

// checks if the BatchInputMarketingEventEmailSubscriber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputMarketingEventEmailSubscriber{}

// BatchInputMarketingEventEmailSubscriber List of marketing event details to create or update
type BatchInputMarketingEventEmailSubscriber struct {
	// List of marketing event details to create or update
	Inputs []MarketingEventEmailSubscriber `json:"inputs"`
}

// NewBatchInputMarketingEventEmailSubscriber instantiates a new BatchInputMarketingEventEmailSubscriber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputMarketingEventEmailSubscriber(inputs []MarketingEventEmailSubscriber) *BatchInputMarketingEventEmailSubscriber {
	this := BatchInputMarketingEventEmailSubscriber{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputMarketingEventEmailSubscriberWithDefaults instantiates a new BatchInputMarketingEventEmailSubscriber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputMarketingEventEmailSubscriberWithDefaults() *BatchInputMarketingEventEmailSubscriber {
	this := BatchInputMarketingEventEmailSubscriber{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputMarketingEventEmailSubscriber) GetInputs() []MarketingEventEmailSubscriber {
	if o == nil {
		var ret []MarketingEventEmailSubscriber
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputMarketingEventEmailSubscriber) GetInputsOk() ([]MarketingEventEmailSubscriber, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputMarketingEventEmailSubscriber) SetInputs(v []MarketingEventEmailSubscriber) {
	o.Inputs = v
}

func (o BatchInputMarketingEventEmailSubscriber) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputMarketingEventEmailSubscriber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputMarketingEventEmailSubscriber struct {
	value *BatchInputMarketingEventEmailSubscriber
	isSet bool
}

func (v NullableBatchInputMarketingEventEmailSubscriber) Get() *BatchInputMarketingEventEmailSubscriber {
	return v.value
}

func (v *NullableBatchInputMarketingEventEmailSubscriber) Set(val *BatchInputMarketingEventEmailSubscriber) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputMarketingEventEmailSubscriber) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputMarketingEventEmailSubscriber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputMarketingEventEmailSubscriber(val *BatchInputMarketingEventEmailSubscriber) *NullableBatchInputMarketingEventEmailSubscriber {
	return &NullableBatchInputMarketingEventEmailSubscriber{value: val, isSet: true}
}

func (v NullableBatchInputMarketingEventEmailSubscriber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputMarketingEventEmailSubscriber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
